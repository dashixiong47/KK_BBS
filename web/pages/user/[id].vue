<template>
  <div class="">
    <UserInfo></UserInfo>
    <div class="pt-5">
      <NuxtPage></NuxtPage>
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
