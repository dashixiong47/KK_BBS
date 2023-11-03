<template>
  <NuxtLayout>
    <NuxtLoadingIndicator class="!transform-none top-16" />
    <NuxtPage />
  </NuxtLayout>
  <Teleport to="body">
    <FloatingList />
  </Teleport>
</template>
<script setup></script>
