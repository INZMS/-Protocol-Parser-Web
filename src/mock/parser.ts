export const mockParseResult = {
  protocol: "JT-808",
  messageId: "0200",
  messageName: "位置上传",
  length: 69,

  fields: [
    { name: "消息ID", offset: 0, len: 2, hex: "0200", value: "位置上传", desc: "消息类型" },
    { name: "消息体属性", offset: 2, len: 2, hex: "0045", value: "定位数据", desc: "属性字段" },
    { name: "终端手机号", offset: 4, len: 6, hex: "013800138000", value: "13800138000", desc: "终端号码" },
    { name: "流水号", offset: 10, len: 2, hex: "0001", value: "1", desc: "消息流水号" },
    { name: "报警标志", offset: 14, len: 4, hex: "C8001B5E", value: "正常", desc: "报警状态" },
    { name: "状态", offset: 18, len: 4, hex: "E3003C18", value: "ACC开,北斗定位", desc: "状态位" },
    { name: "纬度", offset: 22, len: 4, hex: "18071D0C", value: "31.230400°", desc: "纬度(度)" },
    { name: "经度", offset: 26, len: 4, hex: "0A1E0000", value: "121.473700°", desc: "经度(度)" },
    { name: "速度", offset: 30, len: 2, hex: "003C", value: "60 km/h", desc: "单位:0.1km/h" },
    { name: "方向", offset: 32, len: 2, hex: "005A", value: "90°", desc: "单位:0.1°" },
    { name: "定位时间", offset: 34, len: 6, hex: "260729103025", value: "2026-07-29 10:30:25", desc: "YYMMDDhhmmss" }
  ],

  rawHex: "7E020000450100000000013800138000018071D0C0A1E0000003C005A2607291030257E",

  json: {
    protocol: "JT-808",
    messageId: "0200",
    messageName: "位置上传",
    data: {
      terminalNo: "13800138000",
      latitude: 31.2304,
      longitude: 121.4737,
      speed: 60,
      direction: 90,
      time: "2026-07-29 10:30:25"
    }
  }
};

export const mockExamples = [
  { label: "位置上传(0200)", hex: mockParseResult.rawHex },
  { label: "报警信息(0801)", hex: "7E0801000000007E" },
  { label: "终端心跳(0002)", hex: "7E0002000000007E" },
  { label: "终端注册(0100)", hex: "7E010000000000320034303030303030303030307E" },
  { label: "终端注册应答(8100)", hex: "7E810000000000317E" },
  { label: "位置批量上传(0704)", hex: "7E070400000000517E" }
];

export const mockHistory = [
  { id: 1, time: "2026-07-26 10:20:30", protocol: "JT-808", messageId: "0200", messageName: "位置上传", length: 69 },
  { id: 2, time: "2026-07-26 10:15:22", protocol: "JT-808", messageId: "0002", messageName: "终端心跳", length: 13 },
  { id: 3, time: "2026-07-26 10:10:05", protocol: "JT-808", messageId: "0100", messageName: "终端注册", length: 32 },
  { id: 4, time: "2026-07-25 15:43:11", protocol: "2929协议", messageId: "8100", messageName: "终端登录应答", length: 25 },
  { id: 5, time: "2026-07-25 15:40:02", protocol: "2929协议", messageId: "8102", messageName: "位置上传报文", length: 56 },
  { id: 6, time: "2026-07-25 09:12:47", protocol: "JT-808", messageId: "0704", messageName: "位置批量上传", length: 128 },
  { id: 7, time: "2026-07-24 21:03:19", protocol: "JT-808", messageId: "0200", messageName: "位置上传", length: 69 },
  { id: 8, time: "2026-07-24 18:55:40", protocol: "JT-808", messageId: "0002", messageName: "终端心跳", length: 13 },
  { id: 9, time: "2026-07-24 11:30:08", protocol: "2929协议", messageId: "8100", messageName: "终端登录应答", length: 25 },
  { id: 10, time: "2026-07-23 20:47:53", protocol: "JT-808", messageId: "0801", messageName: "报警信息", length: 41 },
  { id: 11, time: "2026-07-23 14:22:16", protocol: "JT-808", messageId: "0100", messageName: "终端注册", length: 32 },
  { id: 12, time: "2026-07-23 09:05:33", protocol: "2929协议", messageId: "8102", messageName: "位置上传报文", length: 56 }
];
