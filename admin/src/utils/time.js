export let formatTime = (dateString) => {
    const options = { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false };
    // 解析日期时间字符串
    const date = new Date(dateString);
    // 使用本地时间和自定义选项进行格式化
    return date.toLocaleDateString('zh-CN', options).replace(/\//g, '-');

}