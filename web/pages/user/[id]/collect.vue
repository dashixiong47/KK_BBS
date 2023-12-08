<template>
  <List :initFunc="1" :userId="route.params.id"></List>
</template>

<script setup>
import List from "~/components/list";
const route = useRoute();
</script>
