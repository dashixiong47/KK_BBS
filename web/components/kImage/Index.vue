<template>
  <div class="bg-slate-400 w-full h-full overflow-hidden" v-lazyload="source" />
</template>

<script setup>
const props = defineProps({
  source: {
    type: String,
    required: "",
  },
});
</script>
