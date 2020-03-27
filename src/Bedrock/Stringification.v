Require Import Coq.ZArith.ZArith.
Require Import Coq.Lists.List.
Require Import Coq.Strings.String.
Require Import bedrock2.Syntax.
Require Import bedrock2.ToCString.
Require Import bedrock2.BasicC64Semantics.
Require Import Crypto.Stringification.Language.
Require Import Crypto.Stringification.IR.
Require Import Crypto.Bedrock.Defaults.
Require Import Crypto.Bedrock.Types.
Require Import Crypto.Bedrock.Translation.Func.
Require Import Crypto.Language.API.
Import Stringification.Language.Compilers.
Import Types.

Import ListNotations Types.Notations.
Local Open Scope string_scope.
Local Open Scope list_scope.

Fixpoint make_names (prefix : string)
         (nextn : nat) (t : base.type) : option (nat * base_ltype t) :=
  match t as t0 return option (nat * base_ltype t0) with
  | base.type.prod a b =>
    match make_names prefix nextn a with
    | Some (nextn, anames) =>
      match make_names prefix nextn b with
      | Some (nextn, bnames) => Some (nextn, (anames, bnames))
      | None => None
      end
    | None => None
    end
  | base_listZ =>
    let num := Decimal.decimal_string_of_Z (Z.of_nat nextn) in
    Some (S nextn, String.append prefix num)
  | base_Z =>
    let num := Decimal.decimal_string_of_Z (Z.of_nat nextn) in
    Some (S nextn, String.append prefix num)
  | _ => None
  end.
Fixpoint make_innames' (nextn : nat) (t : API.type)
  : option (nat * type.for_each_lhs_of_arrow ltype t) :=
  match t with
  | type.base _ => Some (nextn, tt)
  | type.arrow (type.base s) d =>
    match make_names "in" nextn s with
    | Some (nextn, snames) =>
      match make_innames' nextn d with
      | Some (nextn, dnames) =>
        Some (nextn, (snames, dnames))
      | None => None
      end
    | None => None
    end
  | type.arrow _ _ => None
  end.
Definition make_innames t : option (type.for_each_lhs_of_arrow ltype t) :=
  option_map snd (make_innames' 0 t).
Definition make_outnames t : option (base_ltype t) :=
  option_map snd (make_names "out" 0 t).

(* mostly a duplicate of list_lengths_from_value, just with ZRange interp *)
Fixpoint list_lengths_from_bounds {t}
  : ZRange.type.base.option.interp t -> option (base_listonly nat t) :=
  match t as t0 return
        ZRange.type.base.option.interp t0 -> option (base_listonly nat t0) with
  | base.type.prod a b =>
    fun x =>
      Option.bind
        (list_lengths_from_bounds (fst x))
        (fun x1 =>
           option_map
             (fun x2 => (x1, x2))
             (list_lengths_from_bounds (snd x)))
  | base_listZ =>
    fun x : option (list _) => option_map (@List.length _) x
  | _ => fun _ => Some tt
  end.
Fixpoint list_lengths_from_argbounds {t}
  : type.for_each_lhs_of_arrow ZRange.type.option.interp t ->
    option (type.for_each_lhs_of_arrow list_lengths t) :=
  match t as t0 return
        type.for_each_lhs_of_arrow _ t0 ->
        option (type.for_each_lhs_of_arrow _ t0) with
  | type.base b => fun _ => Some tt
  | type.arrow (type.base a) b =>
    fun x =>
      Option.bind
        (list_lengths_from_bounds (fst x))
        (fun x1 =>
           option_map
             (fun x2 => (x1, x2))
             (list_lengths_from_argbounds (snd x)))
  | type.arrow a b => fun _ => None
  end.

Definition bedrock_func_to_lines (f : bedrock_func)
  : list string :=
  [@c_func BasicCSyntax.to_c_parameters f].

(* TODO: for now, name_list is just ignored -- could probably make it not ignored *)
Definition Bedrock2_ToFunctionLines
           {relax_zrange : relax_zrange_opt}
           (do_bounds_check : bool) (static : bool) (prefix : string) (name : string)
           {t}
           (e : @API.Expr t)
           (comment : type.for_each_lhs_of_arrow ToString.OfPHOAS.var_data t ->
                      ToString.OfPHOAS.var_data (type.base (type.final_codomain t)) -> list string)
           (name_list : option (list string))
           (inbounds : type.for_each_lhs_of_arrow ZRange.type.option.interp t)
           (outbounds : ZRange.type.base.option.interp (type.final_codomain t))
  : (list string * ToString.ident_infos) + string
  :=
    match make_innames t, make_outnames (type.final_codomain t),
          list_lengths_from_argbounds inbounds with
    | Some innames, Some outnames, Some lengths =>
      let f := (name, translate_func e innames lengths outnames) in
      if error_free_cmd (snd (snd f))
      then inl (bedrock_func_to_lines f, ToString.ident_info_empty)
      else inr
             (String.concat
                String.NewLine
                ("ERROR-CONTAINING OUTPUT:"
                 :: (bedrock_func_to_lines f)
                 ++ ["Error occured during translation to bedrock2. This is likely because a part of the input expression either had unsupported integer types (bedrock2 requires that all integers have the same size) or contained an unsupported operation."]))
    | None, _, _ =>
      inr ("Error determining argument names")
    | _, None, _ =>
      inr ("Error determining return value names")
    | _, _, None =>
      inr ("Error determining argument lengths")
    end.

Definition OutputBedrock2API : ToString.OutputLanguageAPI :=
  {|
    ToString.comment_block s
    := List.map (fun line => "/* " ++ line ++ " */")%string s;

    ToString.ToFunctionLines := @Bedrock2_ToFunctionLines;

    ToString.header := fun _ _ _ => [];

    ToString.footer := fun _ _ _ => [];

    (** No special handling for any functions *)
    ToString.strip_special_infos infos := infos;

  |}.
