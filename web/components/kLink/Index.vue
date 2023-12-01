<template>
  <NuxtLink :to="localePath(to)" v-if="!isOrigin">
    <slot />
  </NuxtLink>
  <a :href="localePath(to)" v-else>
    <slot />
  </a>
</template>

<script setup>
const localePath = useLocalePath();
let { to } = defineProps({
  to: {
    type: String,
    required: true,
  },
  // 是否是原始a标签跳转
  isOrigin: {
    type: Boolean,
    default: false,
  },
});
</script>
