<template>
  <div class="py-5">
    <KImage
      v-for="(item, index) in detail.images.sort((a, b) => a.order - b.order)"
      :key="index"
      :source="getPath(item.url)"
      :alt="item.name"
    />
  </div>
</template>

<script setup>
const { getPath } = usePath();
const props = defineProps({
  detail: {
    type: Object,
    default: () => {},
  },
});
console.log(props.detail);
</script>
