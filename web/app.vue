<template>
  <NuxtLayout>
    <NuxtLoadingIndicator class="!transform-none top-16" />
    <NuxtPage />
  </NuxtLayout>
  <Teleport to="body">
    <FloatingList />
  </Teleport>
</template>
<script setup>
import { useAppConfigStore } from "~/stores/init";
let app = useAppConfigStore();
app.fetchHost();
</script>
