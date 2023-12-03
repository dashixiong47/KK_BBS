<template>
  <div>
    <TopicList class="col-span-6" :userId="route.params.id"></TopicList>
  </div>
</template>

<script setup>
const route = useRoute();
</script>
