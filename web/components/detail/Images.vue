<template>
  <div class="py-5">
    <KImage
      :id="`image_${item.order}`"
      v-for="(item, index) in detail.images"
      @click="openViewer(index)"
      :key="index"
      :source="getPath(item.url)"
      :alt="item.name"
      :rolling="true"
      @show="scrollToImage(item.order)"
      :fillContainer="false"
    />
  </div>
</template>

<script setup>
const { $imageViewer } = useNuxtApp();
const { getPath } = usePath();
const props = defineProps({
  detail: {
    type: Object,
    default: () => {},
  },
});
const emit = defineEmits(["show"]);
function openViewer(index) {
  $imageViewer.open(props.detail.images, index);
}
function scrollToImage(index) {
  emit("show", index);
}
</script>
