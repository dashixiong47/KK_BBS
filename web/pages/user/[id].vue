<template>
  <div class="">
    <UserInfo></UserInfo>
    <!-- <NuxtPage name="user/1"/> -->
    <!-- <p>id: {{ $route.params.id }}</p> -->
    <div class="grid grid-cols-9 mt-5">
      <ul class="h-full m-1 col-span-6">
        <li
          v-for="item in 10"
          class="w-full p-5 rounded-2xl mb-5 shadow-center dark:bg-dark-2"
        >
          <Topic />
        </li>
      </ul>
      <div class="col-span-3">{{ isMobile }}</div>
    </div>
  </div>
</template>

<script setup>
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();
definePageMeta({
  layout: "user",
});
</script>
