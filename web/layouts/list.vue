<template>
  <!-- 导航 -->
  <PcNav />
  <div class="overflow-y-auto relative flex justify-between p-2 xl:p-5  lg:p-5">
    <LeftBox class="sticky top-0 left-0 w-[20%] hidden lg:block  lg:w-[20%]" />
    <main class="w-full flex lg:w-[75%] xl:w-[60%]">
      <div class="w-full">
        <slot />
      </div>
      <!-- <AdditionalPanel class="w-[30%] hidden sm:block md:w-[35%]" /> -->
    </main>
    <RightBox class="sticky top-0 right-0 w-[20%] hidden xl:flex" />
  </div>
</template>

<script setup>
import { useAppConfigStore } from "~/stores/init";
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
useHead({
  title: appConfig.value.appName,
});
</script>
