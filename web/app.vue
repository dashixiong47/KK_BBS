<template>
  <NuxtLayout>
    <NuxtLoadingIndicator />
    <NuxtPage />
  </NuxtLayout>
  <Teleport to="body">
    <FloatingList />
  </Teleport>
</template>
<script setup></script>
