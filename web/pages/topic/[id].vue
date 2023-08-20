<template>
  <div class="">
    <!-- <p>id: {{ $route.params.id }}</p> -->
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-6">
        <Card>
          <div class="flex">
            <Avatar />
            <div class="w-full ml-2 flex items-center justify-between">
              <div class="flex flex-col items-start">
                <KLink to="#" class="text-md font-bold"> 名称 </KLink>
                <span class="text-xs text-gray-400">发布时间</span>
              </div>
              <div>
                <span class="mr-5 text-light-5 text-xs">
                  <Icon name="icon-kejian" class="mr-1" />{{
                    formatNumber(11111)
                  }}
                </span>
                <span class="text-light-5 text-xs">
                  <Icon name="icon-shouye" class="mr-1" />上海
                </span>
              </div>
            </div>
          </div>
          <div class="border-t mt-5 py-5">
          </div>
        </Card>
        <Card class="mt-5">
          <Comments />
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
