#ifndef _CPP_INT_HELPER_H_
#define _CPP_INT_HELPER_H_

#include <iostream>
#include <assert.h>
#include <boost/multiprecision/gmp.hpp>
#include <boost/multiprecision/cpp_int.hpp>

#ifdef _CSV_INCLUDED_
#include "svdpi.h"
#endif

using namespace std;
using namespace boost::multiprecision;
using boost::multiprecision::cpp_int;

// user-defined fixed width cpp_int datastructures can be declared here
typedef number<cpp_int_backend<4096, 4096, unsigned_magnitude, checked, void> >  checked_uint4096_t;

/*
Any utilities that are hidden or not directly available in cpp_int
are generated by this helper class. This way, the user can use the
native cpp_int operations directly and use this as an helper if need
be
*/

class cpp_int_helper {

public:
   cpp_int_helper() { }

   /*
      GETs the bit-slice of the original bit-vector from cpp_int
      val => input cpp_int which needs to be sliced
      min => min-bit-position of the bit-vector
      max => max-bit-position of the bit-vector

      Templated to allow for different sized variants
   */
   template <class T>
   T get_slc(const T& val, uint32_t min, uint32_t max) {
      assert(min <= max);
      return s_get_slc(val,max-min+1,min);
   }
  // This is in old ac_int size for familiarity.
  template <class T>
  T static s_get_slc (T val, uint32_t sz, uint32_t min) {
      T val_mask = 1;
      val_mask = (val_mask << sz) - 1;
      val = ( (val >> min) & val_mask) ;
      return val;
   }

   /*
      SETs the bit-slice of the original bit-vector from cpp_int
      val => input cpp_int which needs to be sliced
      min => min-bit-position of the bit-vector
      max => max-bit-position of the bit-vector

      Templated to allow for different sized variants
   */

   template <class T>
   const T& set_slc (T&  val, const T& set_val, uint32_t min, uint32_t max) const {
      assert(min <= max);
      /*for (int i=0; i <= (max-min); i++) {
         bit_test(set_val, i) ? bit_set(val,min+i) : bit_unset(val, min+i);
      }
      return val;*/
      ////trying to optimize it, by Changqi
      uint32_t vmsb = 0;
      try { vmsb = msb(val); }
      catch (...) { vmsb = 0; }
      uint32_t m = std::max(vmsb, max);
      cpp_int bm(1);
      bm <<= m;
      bm |= ~bm;
      bm >>= (max + 1);
      bm <<= (max + 1);
      if (min > 0) {
          cpp_int bm2(1);
          bm2 <<= (min - 1);
          bm2 |= ~bm2;
          bm |= bm2;
      }
      val &= bm;

      cpp_int bm2(1);
      bm2 <<= (max - min);
      bm2 |= ~bm2;

      val |= ((set_val & bm2) << min);

      return val;
   }

   template <class V>
   const V& set_slc (V&  val, uint32_t set_val, uint32_t min, uint32_t max) const {
      cpp_int set_val_cpp_int(set_val);

      return set_slc(val, set_val_cpp_int, min, max);
   }

   template <class V>
   static void s_cpp_int_from_array(V & val, uint32_t start, uint32_t end, const uint8_t *ptr) {
     V mask(0);
     if (val != 0) {
       mask = 1;
       mask <<= msb(val)+1;
       mask -= 1;
     }
     V clr_mask = 0xff;
     clr_mask <<= (start*8);
     V tmp;
     for (auto i = start; i <= end; i++) {
       tmp = *ptr;
       if ((mask & clr_mask) != 0) {
	 V effective_mask = (mask ^ clr_mask) & mask;
	 val = (val & effective_mask) | (tmp<<(i*8));
	 clr_mask <<= 8;
       } else { // Value was zero to begin with
	 val = val | (tmp<<(i*8));
       }
       ptr++;
     }
   }

  template <class V, class U>
   static void s_array_from_cpp_int(const V & val, uint32_t start, uint32_t end, uint8_t *ptr, const U & write_strobe) {
     V tmp  = val;
     tmp >>= (start*8);
     for (auto i = start; i <= end; i++) {
       if (bit_test(write_strobe, i)) {
	 *ptr = (tmp & 0xff).template convert_to<uint8_t>();
       }
       tmp >>= 8;
       ptr++;
     }
   }

  template <class V>
   static void s_array_from_cpp_int(const V & val, uint32_t start, uint32_t end, uint8_t *ptr) {
     V tmp  = val;
     tmp >>= (start*8);
     for (auto i = start; i <= end; i++) {
       *ptr = (tmp & 0xff).template convert_to<uint8_t>();
       tmp >>= 8;
       ptr++;
     }
   }

   template <class T>
   void static s_rep_slc (T & val, T set_val, uint32_t min, uint32_t max) {
      assert(min <= max);
      for (int i=0; i <= (max-min); i++) {
         bit_test(set_val, i) ? bit_set(val,min+i) : bit_unset(val, min+i);
      }
   }

   template <class T>
   void static s_rep_slc (T & val, cpp_int set_val, uint32_t min, uint32_t max) {
      assert(min <= max);
      for (int i=0; i <= (max-min); i++) {
         bit_test(set_val, i) ? bit_set(val,min+i) : bit_unset(val, min+i);
      }
   }

   /* Returns the Size of the number of bits used by the cpp_int datastructure (in bits) */
   template <class T>
   uint32_t size (T val) {
      return (val.backend().size()*val.backend().limb_bits);
   }

#ifdef _CSV_INCLUDED_
   void convert_svlogic_to_cpp_int(cpp_int & ret_val, svLogicVecVal * val, int width) {
       for(int ii = 0; ii < (width + 31)/32;ii++) {
           ret_val = set_slc( ret_val , static_cast<cpp_int>(val[ii].aval) , (ii*32), (ii*32) + 32 );
       }
   }

   void convert_cpp_int_to_svlogic(svLogicVecVal * ret_val, cpp_int & val, int width) {
     for(int ii = 0; ii < (width + 31)/32; ii++) {
       ret_val[ii].aval = (s_get_slc(val, 32, (ii*32))).convert_to<uint32_t>();
       ret_val[ii].bval = 0;
     }
   }

   void convert_cpp_int_to_svbit(svBitVecVal * ret_val, cpp_int & val, int width) {
     for(int ii = 0; ii < (width + 31)/32; ii++) {
       ret_val[ii] = (s_get_slc(val, 32, (ii*32))).convert_to<uint32_t>();
     }
   }

   /*
   void convert_vpi_vecval_to_cpp_int(cpp_int & ret_val, p_vpi_vecval * val, int width) {
       for(int ii = 0; ii < (width + 31)/32;ii++) {
           ret_val = set_slc( ret_val , static_cast<cpp_int>(val[ii]->aval) , (ii*32), (ii*32) + 32 );
       }
   }

   void convert_cpp_int_to_vpi_vecval(p_vpi_vecval * ret_val, cpp_int & val, int width) {
     for(int ii = 0; ii < (width + 31)/32; ii++) {
       ret_val[ii]->aval = (s_get_slc(val, 32, (ii*32))).convert_to<uint32_t>();
       ret_val[ii]->bval = 0;
     }
   }
   */

#endif
   };

#endif
