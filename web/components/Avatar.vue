<script setup>
let props = defineProps({
  size: {
    type: Number,
    default: 1,
  },
  url: {
    type: String,
    default: "",
  },
});
const sizeClass = computed(() => {
  switch (props.size) {
    case 1:
      return `w-11 h-11`;

    default:
      return `w-10 h-10`;
  }
});
</script>

<template>
  <div
    class="shadow-center rounded-full bg-slate-200 flex-shrink-0"
    :class="sizeClass"
  >
    <img :src="url" alt="" srcset="" />
  </div>
</template>
