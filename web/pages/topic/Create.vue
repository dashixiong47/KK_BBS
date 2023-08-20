<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-6">
        <Card>
          <div class="text-xl font-bold">发帖子</div>
          <div class="border-t mt-5 py-5">
            <Editor class="mb-5" />
            <Editor />
          </div>
        </Card>
      </div>
      <div class="hidden sm:col-span-3 sm:block m-1">
        <Card class="sticky top-0">
          <p class="border-b pb-2 mb-2 font-bold">附件</p>
          <ul>
            <li>121212</li>
          </ul>
        </Card>
      </div>
    </div>
  </div>
</template>

<script setup>
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();
definePageMeta({
  layout: "user",
});
</script>
