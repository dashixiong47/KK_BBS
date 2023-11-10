<template>
  <button
    @click="throttledClick"
    class="relative px-5 py-2 font-main-color"
    :class="{
      'active:shadow-active': !porps.disabled,
    }"
    :disabled="porps.disabled"
  >
    <div
      v-if="actived"
      class="absolute h-10 w-4 -right-0 -top-3 bg-[--color-primary] dark:bg-[--dark-color-primary] -rotate-45"
    >
      <Icon
        name="ep:check"
        size="0.8rem"
        class="text-[--color-white] dark:text-[--dark-color-white] absolute rotate-45 left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
      />
    </div>
    <slot></slot>
  </button>
</template>

<script setup>
import { ref } from "vue";

const porps = defineProps({
  actived: {
    type: Boolean,
    default: false,
  },
  throttle: {
    type: Function,
    default: null,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
});
const { actived, throttle } = porps;

let lastClicked = ref(0);

const throttleTime = 1000; // 节流时间设为1000毫秒（1秒）

const throttledClick = () => {
  if (porps.disabled || typeof throttle !== "function") return;
  const now = Date.now();
  if (now - lastClicked.value >= throttleTime) {
    // 执行你的点击处理逻辑
    lastClicked.value = now;
    throttle();
  }
};
</script>
