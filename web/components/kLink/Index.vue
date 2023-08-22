<template>
  <NuxtLink :to="localePath(to)">
    <slot />
  </NuxtLink>
</template>

<script setup>
const localePath = useLocalePath();
let { to } = defineProps({
  to: {
    type: String,
    required: true,
  },
});
</script>
