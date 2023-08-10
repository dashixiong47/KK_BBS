<template>
  <!-- 导航 -->
  <PcNav class="z-10" />
  <div
    class="h-full grid grid-cols-24 p-3 sm:p-5 overflow-y-scroll relative z-0"
  >
    <main class="col-span-24 sm:col-start-3 sm:col-end-23">
      <div class="h-full sm:mr-2">
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
