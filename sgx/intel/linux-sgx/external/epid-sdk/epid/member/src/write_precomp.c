/*############################################################################
  # Copyright 2016-2017 Intel Corporation
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  ############################################################################*/
/// EpidMemberWritePrecomp implementation.
/*! \file */
#include "epid/member/src/write_precomp.h"

#include <stddef.h>
#include <stdint.h>

#include "epid/member/src/context.h"

EpidStatus EpidMemberWritePrecomp(MemberCtx const* ctx,
                                  MemberPrecomp* precomp) {
  if (!ctx) {
    return kEpidBadArgErr;
  }
  if (!precomp) {
    return kEpidBadArgErr;
  }

  *precomp = ctx->precomp;
  return kEpidNoErr;
}
