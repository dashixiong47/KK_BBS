// composables/useTime.js
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import 'dayjs/locale/zh-cn';  // 加载简体中文
import 'dayjs/locale/en';  // 加载英文

export default function useTime() {
    const format = (time) => {
        return dayjs(time).format('YYYY-MM-DD HH:mm:ss');
    }
    const getRelativeTime = (time) => {
        const { locale } = useI18n();
        // 使用相对时间插件
        dayjs.extend(relativeTime);

        // 设置语言为简体中文
        dayjs.locale(mapLocale(locale.value));

        // 示例时间，你可以用实际的时间替换这个
        const someTime = dayjs(time);
        return someTime.fromNow()
    }
    const mapLocale = (i18nLocale) => {
        const localeMap = {
          'zh': 'zh-cn',
          'en': 'en',
          // 添加其他映射
        };
        return localeMap[i18nLocale] || i18nLocale;
      };
    return { format,getRelativeTime };
}
