<template>
  <button
    @click="throttledClick"
    class="shadow-center transition-all relative rounded-lg font-main-color overflow-hidden hover:text-[--illuminate-color]"
    :class="{
      'active:shadow-active': !props.disabled,
      'h-6 px-4 text-xs leading-6': size === 'mini',
      'h-8 px-5 text-sm leading-8': size === 'small',
      'h-11 px-6 leading-11': size === 'default',
      'h-14 px-7 text-lg  leading-14': size === 'big',
    }"
    :disabled="props.disabled"
  >
    <div
      v-if="actived"
      class="absolute h-10 w-4 -right-0 -top-3 text-white bg-[--illuminate-color] dark:bg-[--dark-illuminate-color] -rotate-45"
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

const props = defineProps({
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
  size: {
    type: String,
    default: "default",
    validator: (value) => ["mini", "small", "default", "big"].includes(value),
  },
});
const { actived, throttle } = toRefs(props);

let lastClicked = ref(0);

const throttleTime = 1000; // 节流时间设为1000毫秒（1秒）

const throttledClick = () => {
  if (props.disabled || typeof throttle.value !== "function") return;
  const now = Date.now();
  if (now - lastClicked.value >= throttleTime) {
    lastClicked.value = now;
    throttle.value();
  }
};
</script>
