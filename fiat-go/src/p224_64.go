/* Autogenerated: src/ExtractionOCaml/word_by_word_montgomery --lang=Go --no-wide-int --cmovznz-by-mul --widen-carry --widen-bytes p224 '2^224 - 2^96 + 1' 64 mul square add sub opp from_montgomery nonzero selectznz to_bytes from_bytes */
/* curve description: p224 */
/* requested operations: mul, square, add, sub, opp, from_montgomery, nonzero, selectznz, to_bytes, from_bytes */
/* m = 0xffffffffffffffffffffffffffffffff000000000000000000000001 (from "2^224 - 2^96 + 1") */
/* machine_wordsize = 64 (from "64") */
/*                                                                    */
/* NOTE: In addition to the bounds specified above each function, all */
/*   functions synthesized for this Montgomery arithmetic require the */
/*   input to be strictly less than the prime modulus (m), and also   */
/*   require the input to be in the unique saturated representation.  */
/*   All functions also ensure that these two properties are true of  */
/*   return values.                                                   */

package fiat_p224

import "math/bits"


/*
 * The function fiat_p224_cmovznz_u64 is a single-word conditional move.
 * Postconditions:
 *   out1 = (if arg1 = 0 then arg2 else arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xffffffffffffffff]
 *   arg3: [0x0 ~> 0xffffffffffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xffffffffffffffff]
 */
/*inline*/
func fiat_p224_cmovznz_u64(out1 *uint64, arg1 uint64, arg2 uint64, arg3 uint64) {
  var x1 uint64 = (arg1 * 0xffffffffffffffff)
  var x2 uint64 = ((x1 & arg3) | ((^x1) & arg2))
  *out1 = x2
}

/*
 * The function fiat_p224_mul multiplies two field elements in the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 *   0 ≤ eval arg2 < m
 * Postconditions:
 *   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg2)) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 *   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_mul(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
  var x1 uint64 = (arg1[1])
  var x2 uint64 = (arg1[2])
  var x3 uint64 = (arg1[3])
  var x4 uint64 = (arg1[0])
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Mul64(x4, (arg2[3]))
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Mul64(x4, (arg2[2]))
  var x9 uint64
  var x10 uint64
  x9, x10 = bits.Mul64(x4, (arg2[1]))
  var x11 uint64
  var x12 uint64
  x11, x12 = bits.Mul64(x4, (arg2[0]))
  var x13 uint64
  var x14 uint64
  x13, x14 = bits.Add64(x12, x9, 0x0)
  var x15 uint64
  var x16 uint64
  x15, x16 = bits.Add64(x10, x7, x14)
  var x17 uint64
  var x18 uint64
  x17, x18 = bits.Add64(x8, x5, x16)
  var x19 uint64
  x19, _ = bits.Add64(x6, uint64(0x0), x18)
  var x21 uint64
  x21, _ = bits.Mul64(x11, 0xffffffffffffffff)
  var x23 uint64
  var x24 uint64
  x23, x24 = bits.Mul64(x21, 0xffffffff)
  var x25 uint64
  var x26 uint64
  x25, x26 = bits.Mul64(x21, 0xffffffffffffffff)
  var x27 uint64
  var x28 uint64
  x27, x28 = bits.Mul64(x21, 0xffffffff00000000)
  var x29 uint64
  var x30 uint64
  x29, x30 = bits.Add64(x28, x25, 0x0)
  var x31 uint64
  var x32 uint64
  x31, x32 = bits.Add64(x26, x23, x30)
  var x33 uint64
  x33, _ = bits.Add64(x24, uint64(0x0), x32)
  var x36 uint64
  _, x36 = bits.Add64(x11, x21, 0x0)
  var x37 uint64
  var x38 uint64
  x37, x38 = bits.Add64(x13, x27, x36)
  var x39 uint64
  var x40 uint64
  x39, x40 = bits.Add64(x15, x29, x38)
  var x41 uint64
  var x42 uint64
  x41, x42 = bits.Add64(x17, x31, x40)
  var x43 uint64
  var x44 uint64
  x43, x44 = bits.Add64(x19, x33, x42)
  var x45 uint64
  x45, _ = bits.Add64(uint64(0x0), uint64(0x0), x44)
  var x47 uint64
  var x48 uint64
  x47, x48 = bits.Mul64(x1, (arg2[3]))
  var x49 uint64
  var x50 uint64
  x49, x50 = bits.Mul64(x1, (arg2[2]))
  var x51 uint64
  var x52 uint64
  x51, x52 = bits.Mul64(x1, (arg2[1]))
  var x53 uint64
  var x54 uint64
  x53, x54 = bits.Mul64(x1, (arg2[0]))
  var x55 uint64
  var x56 uint64
  x55, x56 = bits.Add64(x54, x51, 0x0)
  var x57 uint64
  var x58 uint64
  x57, x58 = bits.Add64(x52, x49, x56)
  var x59 uint64
  var x60 uint64
  x59, x60 = bits.Add64(x50, x47, x58)
  var x61 uint64
  x61, _ = bits.Add64(x48, uint64(0x0), x60)
  var x63 uint64
  var x64 uint64
  x63, x64 = bits.Add64(x37, x53, 0x0)
  var x65 uint64
  var x66 uint64
  x65, x66 = bits.Add64(x39, x55, x64)
  var x67 uint64
  var x68 uint64
  x67, x68 = bits.Add64(x41, x57, x66)
  var x69 uint64
  var x70 uint64
  x69, x70 = bits.Add64(x43, x59, x68)
  var x71 uint64
  var x72 uint64
  x71, x72 = bits.Add64(x45, x61, x70)
  var x73 uint64
  x73, _ = bits.Mul64(x63, 0xffffffffffffffff)
  var x75 uint64
  var x76 uint64
  x75, x76 = bits.Mul64(x73, 0xffffffff)
  var x77 uint64
  var x78 uint64
  x77, x78 = bits.Mul64(x73, 0xffffffffffffffff)
  var x79 uint64
  var x80 uint64
  x79, x80 = bits.Mul64(x73, 0xffffffff00000000)
  var x81 uint64
  var x82 uint64
  x81, x82 = bits.Add64(x80, x77, 0x0)
  var x83 uint64
  var x84 uint64
  x83, x84 = bits.Add64(x78, x75, x82)
  var x85 uint64
  x85, _ = bits.Add64(x76, uint64(0x0), x84)
  var x88 uint64
  _, x88 = bits.Add64(x63, x73, 0x0)
  var x89 uint64
  var x90 uint64
  x89, x90 = bits.Add64(x65, x79, x88)
  var x91 uint64
  var x92 uint64
  x91, x92 = bits.Add64(x67, x81, x90)
  var x93 uint64
  var x94 uint64
  x93, x94 = bits.Add64(x69, x83, x92)
  var x95 uint64
  var x96 uint64
  x95, x96 = bits.Add64(x71, x85, x94)
  var x97 uint64
  x97, _ = bits.Add64(x72, uint64(0x0), x96)
  var x99 uint64
  var x100 uint64
  x99, x100 = bits.Mul64(x2, (arg2[3]))
  var x101 uint64
  var x102 uint64
  x101, x102 = bits.Mul64(x2, (arg2[2]))
  var x103 uint64
  var x104 uint64
  x103, x104 = bits.Mul64(x2, (arg2[1]))
  var x105 uint64
  var x106 uint64
  x105, x106 = bits.Mul64(x2, (arg2[0]))
  var x107 uint64
  var x108 uint64
  x107, x108 = bits.Add64(x106, x103, 0x0)
  var x109 uint64
  var x110 uint64
  x109, x110 = bits.Add64(x104, x101, x108)
  var x111 uint64
  var x112 uint64
  x111, x112 = bits.Add64(x102, x99, x110)
  var x113 uint64
  x113, _ = bits.Add64(x100, uint64(0x0), x112)
  var x115 uint64
  var x116 uint64
  x115, x116 = bits.Add64(x89, x105, 0x0)
  var x117 uint64
  var x118 uint64
  x117, x118 = bits.Add64(x91, x107, x116)
  var x119 uint64
  var x120 uint64
  x119, x120 = bits.Add64(x93, x109, x118)
  var x121 uint64
  var x122 uint64
  x121, x122 = bits.Add64(x95, x111, x120)
  var x123 uint64
  var x124 uint64
  x123, x124 = bits.Add64(x97, x113, x122)
  var x125 uint64
  x125, _ = bits.Mul64(x115, 0xffffffffffffffff)
  var x127 uint64
  var x128 uint64
  x127, x128 = bits.Mul64(x125, 0xffffffff)
  var x129 uint64
  var x130 uint64
  x129, x130 = bits.Mul64(x125, 0xffffffffffffffff)
  var x131 uint64
  var x132 uint64
  x131, x132 = bits.Mul64(x125, 0xffffffff00000000)
  var x133 uint64
  var x134 uint64
  x133, x134 = bits.Add64(x132, x129, 0x0)
  var x135 uint64
  var x136 uint64
  x135, x136 = bits.Add64(x130, x127, x134)
  var x137 uint64
  x137, _ = bits.Add64(x128, uint64(0x0), x136)
  var x140 uint64
  _, x140 = bits.Add64(x115, x125, 0x0)
  var x141 uint64
  var x142 uint64
  x141, x142 = bits.Add64(x117, x131, x140)
  var x143 uint64
  var x144 uint64
  x143, x144 = bits.Add64(x119, x133, x142)
  var x145 uint64
  var x146 uint64
  x145, x146 = bits.Add64(x121, x135, x144)
  var x147 uint64
  var x148 uint64
  x147, x148 = bits.Add64(x123, x137, x146)
  var x149 uint64
  x149, _ = bits.Add64(x124, uint64(0x0), x148)
  var x151 uint64
  var x152 uint64
  x151, x152 = bits.Mul64(x3, (arg2[3]))
  var x153 uint64
  var x154 uint64
  x153, x154 = bits.Mul64(x3, (arg2[2]))
  var x155 uint64
  var x156 uint64
  x155, x156 = bits.Mul64(x3, (arg2[1]))
  var x157 uint64
  var x158 uint64
  x157, x158 = bits.Mul64(x3, (arg2[0]))
  var x159 uint64
  var x160 uint64
  x159, x160 = bits.Add64(x158, x155, 0x0)
  var x161 uint64
  var x162 uint64
  x161, x162 = bits.Add64(x156, x153, x160)
  var x163 uint64
  var x164 uint64
  x163, x164 = bits.Add64(x154, x151, x162)
  var x165 uint64
  x165, _ = bits.Add64(x152, uint64(0x0), x164)
  var x167 uint64
  var x168 uint64
  x167, x168 = bits.Add64(x141, x157, 0x0)
  var x169 uint64
  var x170 uint64
  x169, x170 = bits.Add64(x143, x159, x168)
  var x171 uint64
  var x172 uint64
  x171, x172 = bits.Add64(x145, x161, x170)
  var x173 uint64
  var x174 uint64
  x173, x174 = bits.Add64(x147, x163, x172)
  var x175 uint64
  var x176 uint64
  x175, x176 = bits.Add64(x149, x165, x174)
  var x177 uint64
  x177, _ = bits.Mul64(x167, 0xffffffffffffffff)
  var x179 uint64
  var x180 uint64
  x179, x180 = bits.Mul64(x177, 0xffffffff)
  var x181 uint64
  var x182 uint64
  x181, x182 = bits.Mul64(x177, 0xffffffffffffffff)
  var x183 uint64
  var x184 uint64
  x183, x184 = bits.Mul64(x177, 0xffffffff00000000)
  var x185 uint64
  var x186 uint64
  x185, x186 = bits.Add64(x184, x181, 0x0)
  var x187 uint64
  var x188 uint64
  x187, x188 = bits.Add64(x182, x179, x186)
  var x189 uint64
  x189, _ = bits.Add64(x180, uint64(0x0), x188)
  var x192 uint64
  _, x192 = bits.Add64(x167, x177, 0x0)
  var x193 uint64
  var x194 uint64
  x193, x194 = bits.Add64(x169, x183, x192)
  var x195 uint64
  var x196 uint64
  x195, x196 = bits.Add64(x171, x185, x194)
  var x197 uint64
  var x198 uint64
  x197, x198 = bits.Add64(x173, x187, x196)
  var x199 uint64
  var x200 uint64
  x199, x200 = bits.Add64(x175, x189, x198)
  var x201 uint64
  x201, _ = bits.Add64(x176, uint64(0x0), x200)
  var x203 uint64
  var x204 uint64
  x203, x204 = bits.Sub64(x193, 0x1, uint64(0x0))
  var x205 uint64
  var x206 uint64
  x205, x206 = bits.Sub64(x195, 0xffffffff00000000, x204)
  var x207 uint64
  var x208 uint64
  x207, x208 = bits.Sub64(x197, 0xffffffffffffffff, x206)
  var x209 uint64
  var x210 uint64
  x209, x210 = bits.Sub64(x199, 0xffffffff, x208)
  var x212 uint64
  _, x212 = bits.Sub64(x201, uint64(0x0), x210)
  var x213 uint64
  fiat_p224_cmovznz_u64(&x213, x212, x203, x193)
  var x214 uint64
  fiat_p224_cmovznz_u64(&x214, x212, x205, x195)
  var x215 uint64
  fiat_p224_cmovznz_u64(&x215, x212, x207, x197)
  var x216 uint64
  fiat_p224_cmovznz_u64(&x216, x212, x209, x199)
  out1[0] = x213
  out1[1] = x214
  out1[2] = x215
  out1[3] = x216
}

/*
 * The function fiat_p224_square squares a field element in the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 * Postconditions:
 *   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg1)) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_square(out1 *[4]uint64, arg1 *[4]uint64) {
  var x1 uint64 = (arg1[1])
  var x2 uint64 = (arg1[2])
  var x3 uint64 = (arg1[3])
  var x4 uint64 = (arg1[0])
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Mul64(x4, (arg1[3]))
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Mul64(x4, (arg1[2]))
  var x9 uint64
  var x10 uint64
  x9, x10 = bits.Mul64(x4, (arg1[1]))
  var x11 uint64
  var x12 uint64
  x11, x12 = bits.Mul64(x4, (arg1[0]))
  var x13 uint64
  var x14 uint64
  x13, x14 = bits.Add64(x12, x9, 0x0)
  var x15 uint64
  var x16 uint64
  x15, x16 = bits.Add64(x10, x7, x14)
  var x17 uint64
  var x18 uint64
  x17, x18 = bits.Add64(x8, x5, x16)
  var x19 uint64
  x19, _ = bits.Add64(x6, uint64(0x0), x18)
  var x21 uint64
  x21, _ = bits.Mul64(x11, 0xffffffffffffffff)
  var x23 uint64
  var x24 uint64
  x23, x24 = bits.Mul64(x21, 0xffffffff)
  var x25 uint64
  var x26 uint64
  x25, x26 = bits.Mul64(x21, 0xffffffffffffffff)
  var x27 uint64
  var x28 uint64
  x27, x28 = bits.Mul64(x21, 0xffffffff00000000)
  var x29 uint64
  var x30 uint64
  x29, x30 = bits.Add64(x28, x25, 0x0)
  var x31 uint64
  var x32 uint64
  x31, x32 = bits.Add64(x26, x23, x30)
  var x33 uint64
  x33, _ = bits.Add64(x24, uint64(0x0), x32)
  var x36 uint64
  _, x36 = bits.Add64(x11, x21, 0x0)
  var x37 uint64
  var x38 uint64
  x37, x38 = bits.Add64(x13, x27, x36)
  var x39 uint64
  var x40 uint64
  x39, x40 = bits.Add64(x15, x29, x38)
  var x41 uint64
  var x42 uint64
  x41, x42 = bits.Add64(x17, x31, x40)
  var x43 uint64
  var x44 uint64
  x43, x44 = bits.Add64(x19, x33, x42)
  var x45 uint64
  x45, _ = bits.Add64(uint64(0x0), uint64(0x0), x44)
  var x47 uint64
  var x48 uint64
  x47, x48 = bits.Mul64(x1, (arg1[3]))
  var x49 uint64
  var x50 uint64
  x49, x50 = bits.Mul64(x1, (arg1[2]))
  var x51 uint64
  var x52 uint64
  x51, x52 = bits.Mul64(x1, (arg1[1]))
  var x53 uint64
  var x54 uint64
  x53, x54 = bits.Mul64(x1, (arg1[0]))
  var x55 uint64
  var x56 uint64
  x55, x56 = bits.Add64(x54, x51, 0x0)
  var x57 uint64
  var x58 uint64
  x57, x58 = bits.Add64(x52, x49, x56)
  var x59 uint64
  var x60 uint64
  x59, x60 = bits.Add64(x50, x47, x58)
  var x61 uint64
  x61, _ = bits.Add64(x48, uint64(0x0), x60)
  var x63 uint64
  var x64 uint64
  x63, x64 = bits.Add64(x37, x53, 0x0)
  var x65 uint64
  var x66 uint64
  x65, x66 = bits.Add64(x39, x55, x64)
  var x67 uint64
  var x68 uint64
  x67, x68 = bits.Add64(x41, x57, x66)
  var x69 uint64
  var x70 uint64
  x69, x70 = bits.Add64(x43, x59, x68)
  var x71 uint64
  var x72 uint64
  x71, x72 = bits.Add64(x45, x61, x70)
  var x73 uint64
  x73, _ = bits.Mul64(x63, 0xffffffffffffffff)
  var x75 uint64
  var x76 uint64
  x75, x76 = bits.Mul64(x73, 0xffffffff)
  var x77 uint64
  var x78 uint64
  x77, x78 = bits.Mul64(x73, 0xffffffffffffffff)
  var x79 uint64
  var x80 uint64
  x79, x80 = bits.Mul64(x73, 0xffffffff00000000)
  var x81 uint64
  var x82 uint64
  x81, x82 = bits.Add64(x80, x77, 0x0)
  var x83 uint64
  var x84 uint64
  x83, x84 = bits.Add64(x78, x75, x82)
  var x85 uint64
  x85, _ = bits.Add64(x76, uint64(0x0), x84)
  var x88 uint64
  _, x88 = bits.Add64(x63, x73, 0x0)
  var x89 uint64
  var x90 uint64
  x89, x90 = bits.Add64(x65, x79, x88)
  var x91 uint64
  var x92 uint64
  x91, x92 = bits.Add64(x67, x81, x90)
  var x93 uint64
  var x94 uint64
  x93, x94 = bits.Add64(x69, x83, x92)
  var x95 uint64
  var x96 uint64
  x95, x96 = bits.Add64(x71, x85, x94)
  var x97 uint64
  x97, _ = bits.Add64(x72, uint64(0x0), x96)
  var x99 uint64
  var x100 uint64
  x99, x100 = bits.Mul64(x2, (arg1[3]))
  var x101 uint64
  var x102 uint64
  x101, x102 = bits.Mul64(x2, (arg1[2]))
  var x103 uint64
  var x104 uint64
  x103, x104 = bits.Mul64(x2, (arg1[1]))
  var x105 uint64
  var x106 uint64
  x105, x106 = bits.Mul64(x2, (arg1[0]))
  var x107 uint64
  var x108 uint64
  x107, x108 = bits.Add64(x106, x103, 0x0)
  var x109 uint64
  var x110 uint64
  x109, x110 = bits.Add64(x104, x101, x108)
  var x111 uint64
  var x112 uint64
  x111, x112 = bits.Add64(x102, x99, x110)
  var x113 uint64
  x113, _ = bits.Add64(x100, uint64(0x0), x112)
  var x115 uint64
  var x116 uint64
  x115, x116 = bits.Add64(x89, x105, 0x0)
  var x117 uint64
  var x118 uint64
  x117, x118 = bits.Add64(x91, x107, x116)
  var x119 uint64
  var x120 uint64
  x119, x120 = bits.Add64(x93, x109, x118)
  var x121 uint64
  var x122 uint64
  x121, x122 = bits.Add64(x95, x111, x120)
  var x123 uint64
  var x124 uint64
  x123, x124 = bits.Add64(x97, x113, x122)
  var x125 uint64
  x125, _ = bits.Mul64(x115, 0xffffffffffffffff)
  var x127 uint64
  var x128 uint64
  x127, x128 = bits.Mul64(x125, 0xffffffff)
  var x129 uint64
  var x130 uint64
  x129, x130 = bits.Mul64(x125, 0xffffffffffffffff)
  var x131 uint64
  var x132 uint64
  x131, x132 = bits.Mul64(x125, 0xffffffff00000000)
  var x133 uint64
  var x134 uint64
  x133, x134 = bits.Add64(x132, x129, 0x0)
  var x135 uint64
  var x136 uint64
  x135, x136 = bits.Add64(x130, x127, x134)
  var x137 uint64
  x137, _ = bits.Add64(x128, uint64(0x0), x136)
  var x140 uint64
  _, x140 = bits.Add64(x115, x125, 0x0)
  var x141 uint64
  var x142 uint64
  x141, x142 = bits.Add64(x117, x131, x140)
  var x143 uint64
  var x144 uint64
  x143, x144 = bits.Add64(x119, x133, x142)
  var x145 uint64
  var x146 uint64
  x145, x146 = bits.Add64(x121, x135, x144)
  var x147 uint64
  var x148 uint64
  x147, x148 = bits.Add64(x123, x137, x146)
  var x149 uint64
  x149, _ = bits.Add64(x124, uint64(0x0), x148)
  var x151 uint64
  var x152 uint64
  x151, x152 = bits.Mul64(x3, (arg1[3]))
  var x153 uint64
  var x154 uint64
  x153, x154 = bits.Mul64(x3, (arg1[2]))
  var x155 uint64
  var x156 uint64
  x155, x156 = bits.Mul64(x3, (arg1[1]))
  var x157 uint64
  var x158 uint64
  x157, x158 = bits.Mul64(x3, (arg1[0]))
  var x159 uint64
  var x160 uint64
  x159, x160 = bits.Add64(x158, x155, 0x0)
  var x161 uint64
  var x162 uint64
  x161, x162 = bits.Add64(x156, x153, x160)
  var x163 uint64
  var x164 uint64
  x163, x164 = bits.Add64(x154, x151, x162)
  var x165 uint64
  x165, _ = bits.Add64(x152, uint64(0x0), x164)
  var x167 uint64
  var x168 uint64
  x167, x168 = bits.Add64(x141, x157, 0x0)
  var x169 uint64
  var x170 uint64
  x169, x170 = bits.Add64(x143, x159, x168)
  var x171 uint64
  var x172 uint64
  x171, x172 = bits.Add64(x145, x161, x170)
  var x173 uint64
  var x174 uint64
  x173, x174 = bits.Add64(x147, x163, x172)
  var x175 uint64
  var x176 uint64
  x175, x176 = bits.Add64(x149, x165, x174)
  var x177 uint64
  x177, _ = bits.Mul64(x167, 0xffffffffffffffff)
  var x179 uint64
  var x180 uint64
  x179, x180 = bits.Mul64(x177, 0xffffffff)
  var x181 uint64
  var x182 uint64
  x181, x182 = bits.Mul64(x177, 0xffffffffffffffff)
  var x183 uint64
  var x184 uint64
  x183, x184 = bits.Mul64(x177, 0xffffffff00000000)
  var x185 uint64
  var x186 uint64
  x185, x186 = bits.Add64(x184, x181, 0x0)
  var x187 uint64
  var x188 uint64
  x187, x188 = bits.Add64(x182, x179, x186)
  var x189 uint64
  x189, _ = bits.Add64(x180, uint64(0x0), x188)
  var x192 uint64
  _, x192 = bits.Add64(x167, x177, 0x0)
  var x193 uint64
  var x194 uint64
  x193, x194 = bits.Add64(x169, x183, x192)
  var x195 uint64
  var x196 uint64
  x195, x196 = bits.Add64(x171, x185, x194)
  var x197 uint64
  var x198 uint64
  x197, x198 = bits.Add64(x173, x187, x196)
  var x199 uint64
  var x200 uint64
  x199, x200 = bits.Add64(x175, x189, x198)
  var x201 uint64
  x201, _ = bits.Add64(x176, uint64(0x0), x200)
  var x203 uint64
  var x204 uint64
  x203, x204 = bits.Sub64(x193, 0x1, uint64(0x0))
  var x205 uint64
  var x206 uint64
  x205, x206 = bits.Sub64(x195, 0xffffffff00000000, x204)
  var x207 uint64
  var x208 uint64
  x207, x208 = bits.Sub64(x197, 0xffffffffffffffff, x206)
  var x209 uint64
  var x210 uint64
  x209, x210 = bits.Sub64(x199, 0xffffffff, x208)
  var x212 uint64
  _, x212 = bits.Sub64(x201, uint64(0x0), x210)
  var x213 uint64
  fiat_p224_cmovznz_u64(&x213, x212, x203, x193)
  var x214 uint64
  fiat_p224_cmovznz_u64(&x214, x212, x205, x195)
  var x215 uint64
  fiat_p224_cmovznz_u64(&x215, x212, x207, x197)
  var x216 uint64
  fiat_p224_cmovznz_u64(&x216, x212, x209, x199)
  out1[0] = x213
  out1[1] = x214
  out1[2] = x215
  out1[3] = x216
}

/*
 * The function fiat_p224_add adds two field elements in the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 *   0 ≤ eval arg2 < m
 * Postconditions:
 *   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) + eval (from_montgomery arg2)) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 *   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_add(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
  var x1 uint64
  var x2 uint64
  x1, x2 = bits.Add64((arg1[0]), (arg2[0]), 0x0)
  var x3 uint64
  var x4 uint64
  x3, x4 = bits.Add64((arg1[1]), (arg2[1]), x2)
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Add64((arg1[2]), (arg2[2]), x4)
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Add64((arg1[3]), (arg2[3]), x6)
  var x9 uint64
  var x10 uint64
  x9, x10 = bits.Sub64(x1, 0x1, uint64(0x0))
  var x11 uint64
  var x12 uint64
  x11, x12 = bits.Sub64(x3, 0xffffffff00000000, x10)
  var x13 uint64
  var x14 uint64
  x13, x14 = bits.Sub64(x5, 0xffffffffffffffff, x12)
  var x15 uint64
  var x16 uint64
  x15, x16 = bits.Sub64(x7, 0xffffffff, x14)
  var x18 uint64
  _, x18 = bits.Sub64(x8, uint64(0x0), x16)
  var x19 uint64
  fiat_p224_cmovznz_u64(&x19, x18, x9, x1)
  var x20 uint64
  fiat_p224_cmovznz_u64(&x20, x18, x11, x3)
  var x21 uint64
  fiat_p224_cmovznz_u64(&x21, x18, x13, x5)
  var x22 uint64
  fiat_p224_cmovznz_u64(&x22, x18, x15, x7)
  out1[0] = x19
  out1[1] = x20
  out1[2] = x21
  out1[3] = x22
}

/*
 * The function fiat_p224_sub subtracts two field elements in the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 *   0 ≤ eval arg2 < m
 * Postconditions:
 *   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) - eval (from_montgomery arg2)) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 *   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_sub(out1 *[4]uint64, arg1 *[4]uint64, arg2 *[4]uint64) {
  var x1 uint64
  var x2 uint64
  x1, x2 = bits.Sub64((arg1[0]), (arg2[0]), 0x0)
  var x3 uint64
  var x4 uint64
  x3, x4 = bits.Sub64((arg1[1]), (arg2[1]), x2)
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Sub64((arg1[2]), (arg2[2]), x4)
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Sub64((arg1[3]), (arg2[3]), x6)
  var x9 uint64
  fiat_p224_cmovznz_u64(&x9, x8, uint64(0x0), 0xffffffffffffffff)
  var x10 uint64
  var x11 uint64
  x10, x11 = bits.Add64(x1, (x9 & 0x1), 0x0)
  var x12 uint64
  var x13 uint64
  x12, x13 = bits.Add64(x3, (x9 & 0xffffffff00000000), x11)
  var x14 uint64
  var x15 uint64
  x14, x15 = bits.Add64(x5, (x9 & 0xffffffffffffffff), x13)
  var x16 uint64
  x16, _ = bits.Add64(x7, (x9 & 0xffffffff), x15)
  out1[0] = x10
  out1[1] = x12
  out1[2] = x14
  out1[3] = x16
}

/*
 * The function fiat_p224_opp negates a field element in the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 * Postconditions:
 *   eval (from_montgomery out1) mod m = -eval (from_montgomery arg1) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_opp(out1 *[4]uint64, arg1 *[4]uint64) {
  var x1 uint64
  var x2 uint64
  x1, x2 = bits.Sub64(uint64(0x0), (arg1[0]), 0x0)
  var x3 uint64
  var x4 uint64
  x3, x4 = bits.Sub64(uint64(0x0), (arg1[1]), x2)
  var x5 uint64
  var x6 uint64
  x5, x6 = bits.Sub64(uint64(0x0), (arg1[2]), x4)
  var x7 uint64
  var x8 uint64
  x7, x8 = bits.Sub64(uint64(0x0), (arg1[3]), x6)
  var x9 uint64
  fiat_p224_cmovznz_u64(&x9, x8, uint64(0x0), 0xffffffffffffffff)
  var x10 uint64
  var x11 uint64
  x10, x11 = bits.Add64(x1, (x9 & 0x1), 0x0)
  var x12 uint64
  var x13 uint64
  x12, x13 = bits.Add64(x3, (x9 & 0xffffffff00000000), x11)
  var x14 uint64
  var x15 uint64
  x14, x15 = bits.Add64(x5, (x9 & 0xffffffffffffffff), x13)
  var x16 uint64
  x16, _ = bits.Add64(x7, (x9 & 0xffffffff), x15)
  out1[0] = x10
  out1[1] = x12
  out1[2] = x14
  out1[3] = x16
}

/*
 * The function fiat_p224_from_montgomery translates a field element out of the Montgomery domain.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * ((2^64)⁻¹ mod m)^4) mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_from_montgomery(out1 *[4]uint64, arg1 *[4]uint64) {
  var x1 uint64 = (arg1[0])
  var x2 uint64
  x2, _ = bits.Mul64(x1, 0xffffffffffffffff)
  var x4 uint64
  var x5 uint64
  x4, x5 = bits.Mul64(x2, 0xffffffff)
  var x6 uint64
  var x7 uint64
  x6, x7 = bits.Mul64(x2, 0xffffffffffffffff)
  var x8 uint64
  var x9 uint64
  x8, x9 = bits.Mul64(x2, 0xffffffff00000000)
  var x10 uint64
  var x11 uint64
  x10, x11 = bits.Add64(x9, x6, 0x0)
  var x12 uint64
  var x13 uint64
  x12, x13 = bits.Add64(x7, x4, x11)
  var x15 uint64
  _, x15 = bits.Add64(x1, x2, 0x0)
  var x16 uint64
  var x17 uint64
  x16, x17 = bits.Add64(uint64(0x0), x8, x15)
  var x18 uint64
  var x19 uint64
  x18, x19 = bits.Add64(uint64(0x0), x10, x17)
  var x20 uint64
  var x21 uint64
  x20, x21 = bits.Add64(uint64(0x0), x12, x19)
  var x22 uint64
  var x23 uint64
  x22, x23 = bits.Add64(x16, (arg1[1]), 0x0)
  var x24 uint64
  var x25 uint64
  x24, x25 = bits.Add64(x18, uint64(0x0), x23)
  var x26 uint64
  var x27 uint64
  x26, x27 = bits.Add64(x20, uint64(0x0), x25)
  var x28 uint64
  x28, _ = bits.Mul64(x22, 0xffffffffffffffff)
  var x30 uint64
  var x31 uint64
  x30, x31 = bits.Mul64(x28, 0xffffffff)
  var x32 uint64
  var x33 uint64
  x32, x33 = bits.Mul64(x28, 0xffffffffffffffff)
  var x34 uint64
  var x35 uint64
  x34, x35 = bits.Mul64(x28, 0xffffffff00000000)
  var x36 uint64
  var x37 uint64
  x36, x37 = bits.Add64(x35, x32, 0x0)
  var x38 uint64
  var x39 uint64
  x38, x39 = bits.Add64(x33, x30, x37)
  var x41 uint64
  _, x41 = bits.Add64(x22, x28, 0x0)
  var x42 uint64
  var x43 uint64
  x42, x43 = bits.Add64(x24, x34, x41)
  var x44 uint64
  var x45 uint64
  x44, x45 = bits.Add64(x26, x36, x43)
  var x46 uint64
  x46, _ = bits.Add64(x5, uint64(0x0), x13)
  var x48 uint64
  x48, _ = bits.Add64(uint64(0x0), x46, x21)
  var x50 uint64
  x50, _ = bits.Add64(x48, uint64(0x0), x27)
  var x52 uint64
  var x53 uint64
  x52, x53 = bits.Add64(x50, x38, x45)
  var x54 uint64
  var x55 uint64
  x54, x55 = bits.Add64(x42, (arg1[2]), 0x0)
  var x56 uint64
  var x57 uint64
  x56, x57 = bits.Add64(x44, uint64(0x0), x55)
  var x58 uint64
  var x59 uint64
  x58, x59 = bits.Add64(x52, uint64(0x0), x57)
  var x60 uint64
  x60, _ = bits.Mul64(x54, 0xffffffffffffffff)
  var x62 uint64
  var x63 uint64
  x62, x63 = bits.Mul64(x60, 0xffffffff)
  var x64 uint64
  var x65 uint64
  x64, x65 = bits.Mul64(x60, 0xffffffffffffffff)
  var x66 uint64
  var x67 uint64
  x66, x67 = bits.Mul64(x60, 0xffffffff00000000)
  var x68 uint64
  var x69 uint64
  x68, x69 = bits.Add64(x67, x64, 0x0)
  var x70 uint64
  var x71 uint64
  x70, x71 = bits.Add64(x65, x62, x69)
  var x73 uint64
  _, x73 = bits.Add64(x54, x60, 0x0)
  var x74 uint64
  var x75 uint64
  x74, x75 = bits.Add64(x56, x66, x73)
  var x76 uint64
  var x77 uint64
  x76, x77 = bits.Add64(x58, x68, x75)
  var x78 uint64
  x78, _ = bits.Add64(x31, uint64(0x0), x39)
  var x80 uint64
  x80, _ = bits.Add64(uint64(0x0), x78, x53)
  var x82 uint64
  x82, _ = bits.Add64(x80, uint64(0x0), x59)
  var x84 uint64
  var x85 uint64
  x84, x85 = bits.Add64(x82, x70, x77)
  var x86 uint64
  var x87 uint64
  x86, x87 = bits.Add64(x74, (arg1[3]), 0x0)
  var x88 uint64
  var x89 uint64
  x88, x89 = bits.Add64(x76, uint64(0x0), x87)
  var x90 uint64
  var x91 uint64
  x90, x91 = bits.Add64(x84, uint64(0x0), x89)
  var x92 uint64
  x92, _ = bits.Mul64(x86, 0xffffffffffffffff)
  var x94 uint64
  var x95 uint64
  x94, x95 = bits.Mul64(x92, 0xffffffff)
  var x96 uint64
  var x97 uint64
  x96, x97 = bits.Mul64(x92, 0xffffffffffffffff)
  var x98 uint64
  var x99 uint64
  x98, x99 = bits.Mul64(x92, 0xffffffff00000000)
  var x100 uint64
  var x101 uint64
  x100, x101 = bits.Add64(x99, x96, 0x0)
  var x102 uint64
  var x103 uint64
  x102, x103 = bits.Add64(x97, x94, x101)
  var x105 uint64
  _, x105 = bits.Add64(x86, x92, 0x0)
  var x106 uint64
  var x107 uint64
  x106, x107 = bits.Add64(x88, x98, x105)
  var x108 uint64
  var x109 uint64
  x108, x109 = bits.Add64(x90, x100, x107)
  var x110 uint64
  x110, _ = bits.Add64(x63, uint64(0x0), x71)
  var x112 uint64
  x112, _ = bits.Add64(uint64(0x0), x110, x85)
  var x114 uint64
  x114, _ = bits.Add64(x112, uint64(0x0), x91)
  var x116 uint64
  var x117 uint64
  x116, x117 = bits.Add64(x114, x102, x109)
  var x118 uint64
  x118, _ = bits.Add64(x95, uint64(0x0), x103)
  var x120 uint64
  x120, _ = bits.Add64(uint64(0x0), x118, x117)
  var x122 uint64
  var x123 uint64
  x122, x123 = bits.Sub64(x106, 0x1, uint64(0x0))
  var x124 uint64
  var x125 uint64
  x124, x125 = bits.Sub64(x108, 0xffffffff00000000, x123)
  var x126 uint64
  var x127 uint64
  x126, x127 = bits.Sub64(x116, 0xffffffffffffffff, x125)
  var x128 uint64
  var x129 uint64
  x128, x129 = bits.Sub64(x120, 0xffffffff, x127)
  var x131 uint64
  _, x131 = bits.Sub64(uint64(0x0), uint64(0x0), x129)
  var x132 uint64
  fiat_p224_cmovznz_u64(&x132, x131, x122, x106)
  var x133 uint64
  fiat_p224_cmovznz_u64(&x133, x131, x124, x108)
  var x134 uint64
  fiat_p224_cmovznz_u64(&x134, x131, x126, x116)
  var x135 uint64
  fiat_p224_cmovznz_u64(&x135, x131, x128, x120)
  out1[0] = x132
  out1[1] = x133
  out1[2] = x134
  out1[3] = x135
}

/*
 * The function fiat_p224_nonzero outputs a single non-zero word if the input is non-zero and zero otherwise.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 * Postconditions:
 *   out1 = 0 ↔ eval (from_montgomery arg1) mod m = 0
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [0x0 ~> 0xffffffffffffffff]
 */
/*inline*/
func fiat_p224_nonzero(out1 *uint64, arg1 *[4]uint64) {
  var x1 uint64 = ((arg1[0]) | ((arg1[1]) | ((arg1[2]) | ((arg1[3]) | uint64(0x0)))))
  *out1 = x1
}

/*
 * The function fiat_p224_selectznz is a multi-limb conditional select.
 * Postconditions:
 *   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 *   arg3: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff]]
 */
/*inline*/
func fiat_p224_selectznz(out1 *[4]uint64, arg1 uint64, arg2 *[4]uint64, arg3 *[4]uint64) {
  var x1 uint64
  fiat_p224_cmovznz_u64(&x1, arg1, (arg2[0]), (arg3[0]))
  var x2 uint64
  fiat_p224_cmovznz_u64(&x2, arg1, (arg2[1]), (arg3[1]))
  var x3 uint64
  fiat_p224_cmovznz_u64(&x3, arg1, (arg2[2]), (arg3[2]))
  var x4 uint64
  fiat_p224_cmovznz_u64(&x4, arg1, (arg2[3]), (arg3[3]))
  out1[0] = x1
  out1[1] = x2
  out1[2] = x3
  out1[3] = x4
}

/*
 * The function fiat_p224_to_bytes serializes a field element in the Montgomery domain to bytes in little-endian order.
 * Preconditions:
 *   0 ≤ eval arg1 < m
 * Postconditions:
 *   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..31]
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x0], [0x0 ~> 0x0], [0x0 ~> 0x0], [0x0 ~> 0x0]]
 */
/*inline*/
func fiat_p224_to_bytes(out1 *[32]uint64, arg1 *[4]uint64) {
  var x1 uint64 = (arg1[3])
  var x2 uint64 = (arg1[2])
  var x3 uint64 = (arg1[1])
  var x4 uint64 = (arg1[0])
  var x5 uint64 = (x4 >> 8)
  var x6 uint64 = (x4 & 0xff)
  var x7 uint64 = (x5 >> 8)
  var x8 uint64 = (x5 & 0xff)
  var x9 uint64 = (x7 >> 8)
  var x10 uint64 = (x7 & 0xff)
  var x11 uint64 = (x9 >> 8)
  var x12 uint64 = (x9 & 0xff)
  var x13 uint64 = (x11 >> 8)
  var x14 uint64 = (x11 & 0xff)
  var x15 uint64 = (x13 >> 8)
  var x16 uint64 = (x13 & 0xff)
  var x17 uint64 = (x15 >> 8)
  var x18 uint64 = (x15 & 0xff)
  var x19 uint64 = (x17 & 0xff)
  var x20 uint64 = (x3 >> 8)
  var x21 uint64 = (x3 & 0xff)
  var x22 uint64 = (x20 >> 8)
  var x23 uint64 = (x20 & 0xff)
  var x24 uint64 = (x22 >> 8)
  var x25 uint64 = (x22 & 0xff)
  var x26 uint64 = (x24 >> 8)
  var x27 uint64 = (x24 & 0xff)
  var x28 uint64 = (x26 >> 8)
  var x29 uint64 = (x26 & 0xff)
  var x30 uint64 = (x28 >> 8)
  var x31 uint64 = (x28 & 0xff)
  var x32 uint64 = (x30 >> 8)
  var x33 uint64 = (x30 & 0xff)
  var x34 uint64 = (x32 & 0xff)
  var x35 uint64 = (x2 >> 8)
  var x36 uint64 = (x2 & 0xff)
  var x37 uint64 = (x35 >> 8)
  var x38 uint64 = (x35 & 0xff)
  var x39 uint64 = (x37 >> 8)
  var x40 uint64 = (x37 & 0xff)
  var x41 uint64 = (x39 >> 8)
  var x42 uint64 = (x39 & 0xff)
  var x43 uint64 = (x41 >> 8)
  var x44 uint64 = (x41 & 0xff)
  var x45 uint64 = (x43 >> 8)
  var x46 uint64 = (x43 & 0xff)
  var x47 uint64 = (x45 >> 8)
  var x48 uint64 = (x45 & 0xff)
  var x49 uint64 = (x47 & 0xff)
  var x50 uint64 = (x1 >> 8)
  var x51 uint64 = (x1 & 0xff)
  var x52 uint64 = (x50 >> 8)
  var x53 uint64 = (x50 & 0xff)
  var x54 uint64 = (x52 >> 8)
  var x55 uint64 = (x52 & 0xff)
  var x56 uint64 = (x54 & 0xff)
  out1[0] = x6
  out1[1] = x8
  out1[2] = x10
  out1[3] = x12
  out1[4] = x14
  out1[5] = x16
  out1[6] = x18
  out1[7] = x19
  out1[8] = x21
  out1[9] = x23
  out1[10] = x25
  out1[11] = x27
  out1[12] = x29
  out1[13] = x31
  out1[14] = x33
  out1[15] = x34
  out1[16] = x36
  out1[17] = x38
  out1[18] = x40
  out1[19] = x42
  out1[20] = x44
  out1[21] = x46
  out1[22] = x48
  out1[23] = x49
  out1[24] = x51
  out1[25] = x53
  out1[26] = x55
  out1[27] = x56
  out1[28] = uint64(0x0)
  out1[29] = uint64(0x0)
  out1[30] = uint64(0x0)
  out1[31] = uint64(0x0)
}

/*
 * The function fiat_p224_from_bytes deserializes a field element in the Montgomery domain from bytes in little-endian order.
 * Preconditions:
 *   0 ≤ bytes_eval arg1 < m
 * Postconditions:
 *   eval out1 mod m = bytes_eval arg1 mod m
 *   0 ≤ eval out1 < m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x0], [0x0 ~> 0x0], [0x0 ~> 0x0], [0x0 ~> 0x0]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffffffffffff], [0x0 ~> 0xffffffff]]
 */
/*inline*/
func fiat_p224_from_bytes(out1 *[4]uint64, arg1 *[32]uint64) {
  var x1 uint64 = ((arg1[27]) << 24)
  var x2 uint64 = ((arg1[26]) << 16)
  var x3 uint64 = ((arg1[25]) << 8)
  var x4 uint64 = (arg1[24])
  var x5 uint64 = ((arg1[23]) << 56)
  var x6 uint64 = ((arg1[22]) << 48)
  var x7 uint64 = ((arg1[21]) << 40)
  var x8 uint64 = ((arg1[20]) << 32)
  var x9 uint64 = ((arg1[19]) << 24)
  var x10 uint64 = ((arg1[18]) << 16)
  var x11 uint64 = ((arg1[17]) << 8)
  var x12 uint64 = (arg1[16])
  var x13 uint64 = ((arg1[15]) << 56)
  var x14 uint64 = ((arg1[14]) << 48)
  var x15 uint64 = ((arg1[13]) << 40)
  var x16 uint64 = ((arg1[12]) << 32)
  var x17 uint64 = ((arg1[11]) << 24)
  var x18 uint64 = ((arg1[10]) << 16)
  var x19 uint64 = ((arg1[9]) << 8)
  var x20 uint64 = (arg1[8])
  var x21 uint64 = ((arg1[7]) << 56)
  var x22 uint64 = ((arg1[6]) << 48)
  var x23 uint64 = ((arg1[5]) << 40)
  var x24 uint64 = ((arg1[4]) << 32)
  var x25 uint64 = ((arg1[3]) << 24)
  var x26 uint64 = ((arg1[2]) << 16)
  var x27 uint64 = ((arg1[1]) << 8)
  var x28 uint64 = (arg1[0])
  var x29 uint64 = (x28 + (x27 + (x26 + (x25 + (x24 + (x23 + (x22 + x21)))))))
  var x30 uint64 = (x29 & 0xffffffffffffffff)
  var x31 uint64 = (x4 + (x3 + (x2 + x1)))
  var x32 uint64 = (x12 + (x11 + (x10 + (x9 + (x8 + (x7 + (x6 + x5)))))))
  var x33 uint64 = (x20 + (x19 + (x18 + (x17 + (x16 + (x15 + (x14 + x13)))))))
  var x34 uint64 = (x33 & 0xffffffffffffffff)
  var x35 uint64 = (x32 & 0xffffffffffffffff)
  out1[0] = x30
  out1[1] = x34
  out1[2] = x35
  out1[3] = x31
}

