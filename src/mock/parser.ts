export const mockParseResult = {


  protocol:"JT-808",


  messageId:"0200",


  messageName:"位置信息上传",


  length:69,


  fields:[

    {
      name:"消息ID",
      hex:"0200",
      value:"位置上传",
      desc:"消息类型"
    },


    {
      name:"终端手机号",
      hex:"013800138000",
      value:"13800138000",
      desc:"终端号码"
    },


    {
      name:"纬度",
      hex:"18071D0C",
      value:"31.230400",
      desc:"GPS纬度"
    },


    {
      name:"经度",
      hex:"0A1E0000",
      value:"121.473700",
      desc:"GPS经度"
    },


    {
      name:"速度",
      hex:"003C",
      value:"60 km/h",
      desc:"车辆速度"
    },


    {
      name:"方向",
      hex:"005A",
      value:"90°",
      desc:"车辆方向"
    },


    {
      name:"时间",
      hex:"260729103025",
      value:"2026-07-29 10:30:25",
      desc:"GPS时间"
    }

  ],



  rawHex:

`7E020000450100000000013800138000
18071D0C0A1E0000003C005A260729103025
7E`,



  json:{


    protocol:"JT-808",


    messageId:"0200",


    messageName:"位置信息上传",


    data:{


      terminalNo:"13800138000",


      latitude:31.230400,


      longitude:121.473700,


      speed:60,


      direction:90,


      time:"2026-07-29 10:30:25"


    }

  }


}



export const mockHistory=[


{

id:1,

time:"2026-07-29 10:30:25",

protocol:"JT-808",

messageId:"0200",

messageName:"位置信息上传",

length:"69 Bytes"

},


{

id:2,

time:"2026-07-29 10:31:02",

protocol:"JT-808",

messageId:"0002",

messageName:"终端心跳",

length:"12 Bytes"

},


{

id:3,

time:"2026-07-29 10:35:18",

protocol:"2929",

messageId:"0900",

messageName:"车辆登录",

length:"45 Bytes"

}


]