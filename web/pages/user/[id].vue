<template>
  <div class="">
    <UserInfo></UserInfo>
    <div class="pt-5 grid grid-cols-12 gap-5">
      <CardUserCard class="col-span-3 h-96 sticky top-0"> </CardUserCard>
      <div class="col-span-9">
        <NuxtPage></NuxtPage>
      </div>
    </div>
  </div>
</template>

<script setup>
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const route = useRoute();
const { formatNumber } = useFormatNumber();

const { isMobile } = useMobileDetect();
</script>
