<template>
  <PlayVideo class="w-full rounded-xl" :src="url ? getPath(url) : ''" />
</template>

<script setup>
const { getPath } = usePath();
const props = defineProps({
  url: {
    type: String,
    default: "",
  },
});
</script>
