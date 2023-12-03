<template>
  <div class="">
    <UserInfo></UserInfo>
    <!-- <NuxtPage name="user/1"/> -->
    <!-- <p>id: {{ $route.params.id }}</p> -->
    <div class="grid grid-cols-9 mt-5 gap-5">
      
      <NuxtPage class="col-span-6"></NuxtPage>
      <!-- <ul class="h-full  ">
        <li
          v-for="item in 10"
          class="w-full p-5 rounded-2xl mb-5 dark:bg-dark-2"
        >
          <Topic />
        </li>
      </ul> -->
      <div class="col-span-3">
        <Card>
          {{ isMobile }}
        </Card>
      </div>
    </div>
  </div>
</template>

<script setup>
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const route = useRoute();
const { formatNumber } = useFormatNumber();

const { isMobile } = useMobileDetect();

</script>
