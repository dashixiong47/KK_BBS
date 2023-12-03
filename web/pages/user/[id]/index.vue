<template>
  <div class="flex flex-col items-end">
    <KSelect class="mb-4"></KSelect>
    <TopicList class="col-span-6" :userId="route.params.id"></TopicList>
  </div>
</template>

<script setup>
const route = useRoute();
</script>
