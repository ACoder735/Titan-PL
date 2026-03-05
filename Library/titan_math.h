#ifndef TITAN_MATH_H
#define TITAN_MATH_H

#include <stdlib.h>
#include <math.h>

// ===== CONSTANTS =====
TitanVar f_math_pi(void) { return titan_num(3.14159265358979); }
TitanVar f_math_e(void) { return titan_num(2.71828182845904); }
TitanVar f_math_tau(void) { return titan_num(6.28318530717958); }
TitanVar f_math_phi(void) { return titan_num(1.61803398874989); }
TitanVar f_math_sqrt2(void) { return titan_num(1.41421356237309); }
TitanVar f_math_sqrt3(void) { return titan_num(1.73205080756887); }
TitanVar f_math_ln2(void) { return titan_num(0.69314718055994); }
TitanVar f_math_ln10(void) { return titan_num(2.30258509299404); }

// ===== BASIC =====
TitanVar f_math_abs(TitanVar v) { return titan_num(fabs(v.num)); }
TitanVar f_math_fabs(TitanVar v) { return titan_num(fabs(v.num)); }
TitanVar f_math_sqrt(TitanVar v) { return titan_num(sqrt(v.num)); }
TitanVar f_math_cbrt(TitanVar v) { return titan_num(cbrt(v.num)); }
TitanVar f_math_pow(TitanVar a, TitanVar b) { return titan_num(pow(a.num, b.num)); }
TitanVar f_math_exp(TitanVar v) { return titan_num(exp(v.num)); }
TitanVar f_math_exp2(TitanVar v) { return titan_num(exp2(v.num)); }
TitanVar f_math_expm1(TitanVar v) { return titan_num(expm1(v.num)); }
TitanVar f_math_hypot(TitanVar a, TitanVar b) { return titan_num(hypot(a.num, b.num)); }
TitanVar f_math_fma(TitanVar a, TitanVar b, TitanVar c) { return titan_num(fma(a.num, b.num, c.num)); }

// ===== LOGARITHMS =====
TitanVar f_math_log(TitanVar v) { return titan_num(log(v.num)); }
TitanVar f_math_log10(TitanVar v) { return titan_num(log10(v.num)); }
TitanVar f_math_log2(TitanVar v) { return titan_num(log2(v.num)); }
TitanVar f_math_log1p(TitanVar v) { return titan_num(log1p(v.num)); }
TitanVar f_math_logb(TitanVar v) { return titan_num(logb(v.num)); }

// ===== TRIGONOMETRIC =====
TitanVar f_math_sin(TitanVar v) { return titan_num(sin(v.num)); }
TitanVar f_math_cos(TitanVar v) { return titan_num(cos(v.num)); }
TitanVar f_math_tan(TitanVar v) { return titan_num(tan(v.num)); }
TitanVar f_math_asin(TitanVar v) { return titan_num(asin(v.num)); }
TitanVar f_math_acos(TitanVar v) { return titan_num(acos(v.num)); }
TitanVar f_math_atan(TitanVar v) { return titan_num(atan(v.num)); }
TitanVar f_math_atan2(TitanVar y, TitanVar x) { return titan_num(atan2(y.num, x.num)); }

// ===== HYPERBOLIC =====
TitanVar f_math_sinh(TitanVar v) { return titan_num(sinh(v.num)); }
TitanVar f_math_cosh(TitanVar v) { return titan_num(cosh(v.num)); }
TitanVar f_math_tanh(TitanVar v) { return titan_num(tanh(v.num)); }
TitanVar f_math_asinh(TitanVar v) { return titan_num(asinh(v.num)); }
TitanVar f_math_acosh(TitanVar v) { return titan_num(acosh(v.num)); }
TitanVar f_math_atanh(TitanVar v) { return titan_num(atanh(v.num)); }

// ===== ROUNDING =====
TitanVar f_math_floor(TitanVar v) { return titan_num(floor(v.num)); }
TitanVar f_math_ceil(TitanVar v) { return titan_num(ceil(v.num)); }
TitanVar f_math_round(TitanVar v) { return titan_num(round(v.num)); }
TitanVar f_math_trunc(TitanVar v) { return titan_num(trunc(v.num)); }

// ===== REMAINDER =====
TitanVar f_math_fmod(TitanVar a, TitanVar b) { return titan_num(fmod(a.num, b.num)); }
TitanVar f_math_remainder(TitanVar a, TitanVar b) { return titan_num(remainder(a.num, b.num)); }

// ===== MIN/MAX =====
TitanVar f_math_min(TitanVar a, TitanVar b) { return titan_num(a.num < b.num ? a.num : b.num); }
TitanVar f_math_max(TitanVar a, TitanVar b) { return titan_num(a.num > b.num ? a.num : b.num); }
TitanVar f_math_fdim(TitanVar a, TitanVar b) { return titan_num(fdim(a.num, b.num)); }

// ===== SIGN =====
TitanVar f_math_sign(TitanVar v) { 
    if (v.num > 0) return titan_num(1); 
    if (v.num < 0) return titan_num(-1); 
    return titan_num(0); 
}
TitanVar f_math_copysign(TitanVar a, TitanVar b) { return titan_num(copysign(a.num, b.num)); }

// ===== SPECIAL =====
TitanVar f_math_erf(TitanVar v) { return titan_num(erf(v.num)); }
TitanVar f_math_erfc(TitanVar v) { return titan_num(erfc(v.num)); }
TitanVar f_math_lgamma(TitanVar v) { return titan_num(lgamma(v.num)); }
TitanVar f_math_tgamma(TitanVar v) { return titan_num(tgamma(v.num)); }

// ===== BESSEL =====
TitanVar f_math_j0(TitanVar v) { return titan_num(j0(v.num)); }
TitanVar f_math_j1(TitanVar v) { return titan_num(j1(v.num)); }
TitanVar f_math_y0(TitanVar v) { return titan_num(y0(v.num)); }
TitanVar f_math_y1(TitanVar v) { return titan_num(y1(v.num)); }

// ===== CONVERSION =====
TitanVar f_math_degrees(TitanVar radians) { 
    return titan_num(radians.num * 180.0 / 3.14159265358979); 
}
TitanVar f_math_radians(TitanVar degrees) { 
    return titan_num(degrees.num * 3.14159265358979 / 180.0); 
}

// ===== CLAMP =====
TitanVar f_math_clamp(TitanVar v, TitanVar min, TitanVar max) {
    if (v.num < min.num) return min;
    if (v.num > max.num) return max;
    return v;
}

// ===== SUM/PRODUCT/MEAN =====
TitanVar f_math_sum(TitanVar arr) {
    if (arr.type != 3) return titan_num(0);
    double total = 0;
    for (int i = 0; i < arr.arr->count; i++) {
        if (arr.arr->data[i].type == 1) total += arr.arr->data[i].num;
    }
    return titan_num(total);
}

TitanVar f_math_product(TitanVar arr) {
    if (arr.type != 3) return titan_num(1);
    double total = 1;
    for (int i = 0; i < arr.arr->count; i++) {
        if (arr.arr->data[i].type == 1) total *= arr.arr->data[i].num;
    }
    return titan_num(total);
}

TitanVar f_math_mean(TitanVar arr) {
    if (arr.type != 3 || arr.arr->count == 0) return titan_num(0);
    double total = 0;
    for (int i = 0; i < arr.arr->count; i++) {
        if (arr.arr->data[i].type == 1) total += arr.arr->data[i].num;
    }
    return titan_num(total / arr.arr->count);
}

#endif