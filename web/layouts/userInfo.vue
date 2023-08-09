<template>
  <!-- 导航 -->
  <PcNav class="z-10" />
  <div
    class="h-full grid grid-cols-24 p-3 sm:p-5 overflow-y-scroll relative z-0"
  >
    <main
      class="mx-auto col-start-1 col-span-24 xl:col-start-6 xl:col-end-20 grid grid-cols-3"
    >
      <div class="h-full sm:mr-2 col-span-3 sm:col-span-2">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
import { useCounterStore } from "~/stores/main.js";
// 获取pinia实例
const store = useCounterStore();
const { count } = store;
console.log(count);
</script>
