<script setup>
const colorMode = useColorMode();
</script>
<template>
  <div class="test">cs测试</div>
  <p class="p-4 pb-2">
    <select
      v-model="colorMode.preference"
      class="border w-24 h-8 dark:bg-gray-900 dark:text-white dark:border-gray-700"
    >
      <option value="system">System</option>
      <option value="light">Light</option>
      <option value="dark">Dark</option>
    </select>
  </p>
</template>
<style lang="postcss"></style>
