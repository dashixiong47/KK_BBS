<script setup>
let prop = defineProps({
  name: String,
  size: {
    type: String,
    default: "text-base",
  },
});
</script>
<template>
  <i :class="`iconfont ${prop.name} ${prop.size}`"></i>
</template>
