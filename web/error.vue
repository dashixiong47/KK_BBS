<template>
  <button @click="handleError">Clear errors</button>
</template>

<script setup>
definePageMeta({
  layout: "user",
});
const props = defineProps({
  error: Object,
});

const handleError = () => clearError({ redirect: "/" });
</script>
