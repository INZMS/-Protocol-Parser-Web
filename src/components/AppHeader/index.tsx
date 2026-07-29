import {
 QuestionCircleOutlined,
 DeleteOutlined,
 CodeOutlined
} from "@ant-design/icons";


export default function Header(){


return (

<div

style={{

height:80,

background:"#fff",

borderBottom:"1px solid #e5e7eb",

display:"flex",

alignItems:"center",

justifyContent:"space-between",

padding:"0 32px"

}}

>


{/* 左侧 */}

<div

style={{

display:"flex",

alignItems:"center",

gap:14

}}

>


<div

style={{

width:42,

height:42,

borderRadius:8,

background:"#1677ff",

display:"flex",

alignItems:"center",

justifyContent:"center",

color:"#fff",

fontSize:22

}}

>


<CodeOutlined />


</div>



<div>


<div

style={{

fontSize:20,

fontWeight:600,

lineHeight:"28px"

}}

>

协议解析工具

</div>



<div

style={{

fontSize:12,

color:"#8c8c8c",

lineHeight:"18px"

}}

>

Protocol Parser Tool

</div>


</div>


</div>




{/* 右侧 */}


<div

style={{

display:"flex",

alignItems:"center",

gap:24

}}

>


<div>

<QuestionCircleOutlined />

<span style={{marginLeft:6}}>
使用说明
</span>

</div>



<div>

<DeleteOutlined />

<span style={{marginLeft:6}}>
清空记录
</span>

</div>




<div

style={{

display:"flex",

alignItems:"center",

gap:8

}}

>


<div

style={{

width:32,

height:32,

borderRadius:"50%",

background:"#e6f4ff",

display:"flex",

alignItems:"center",

justifyContent:"center"

}}

>

A

</div>


admin


</div>


</div>



</div>

)

}