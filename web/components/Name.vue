<template>
  <span class="cursor-pointer">{{ name }}</span>
</template>

<script setup>
let { name } = defineProps({
  name: {
    type: String,
    default: "name",
  },
});
</script>
