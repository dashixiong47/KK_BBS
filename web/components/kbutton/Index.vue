<template>
  <button
    class="glass relative px-5 py-2 rounded-lg primary-text active:shadow-active"
  >
    <div v-if="actived" class="absolute h-10 w-4 -right-0 -top-3 bg-[--color-primary] dark:bg-[--dark-color-primary] -rotate-45">
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
const { actived } = defineProps({
  actived: {
    type: Boolean,
    default: false,
  },
});
</script>
