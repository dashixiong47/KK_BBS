<template>
  <div class="bg-slate-400 rounded-2xl overflow-hidden" v-lazyload="source" />
</template>

<script setup>
const { source } = defineProps({
  source: {
    type: String,
    required: true,
  },
});
</script>
