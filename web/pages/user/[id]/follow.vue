<template>
  <div>
    <List class="col-span-6" :userId="route.params.id"></List>
  </div>
</template>

<script setup>
import { useAppConfigStore } from "~/stores/init";
const route = useRoute();
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
useHead({
  title: "关注 - " + appConfig.value.appName,
});
</script>
