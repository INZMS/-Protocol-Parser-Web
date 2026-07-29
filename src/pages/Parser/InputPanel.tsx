import {
    Card,
    Select,
    Input,
    Button,
    Space,
    Tag
} from "antd";


import {
    PlayCircleOutlined,
    ClearOutlined
} from "@ant-design/icons";


import {
    useParserStore
} from "../../store/parser";


import {
    mockParseResult
} from "../../mock/parser";




export default function InputPanel() {


    const {

        hex,

        setHex,

        parse,

        clear


    } = useParserStore();




    return (


        <Card

            title="输入报文"

            style={{


                height: "100%",

                overflow: "hidden"

            }}
            styles={{

                body: {

                    height: "calc(100% - 57px)",

                    display: "flex",

                    flexDirection: "column",

                    overflow: "hidden"

                }

            }}

        >



            {/* 协议选择 */}


            <div

                style={{

                    marginBottom: 16

                }}

            >



                <div

                    style={{

                        marginBottom: 8,

                        fontWeight: 500

                    }}

                >

                    协议选择

                </div>




                <Select

                    style={{

                        width: "100%"

                    }}

                    defaultValue="JT808"


                    options={[


                        {

                            label: "JT-808",

                            value: "JT808"

                        },


                        {

                            label: "2929协议",

                            value: "2929"

                        }


                    ]}


                />



            </div>






            {/* HEX输入 */}



            <div>



                <div

                    style={{

                        display: "flex",

                        justifyContent: "space-between",

                        marginBottom: 8

                    }}

                >



                    <span

                        style={{

                            fontWeight: 500

                        }}

                    >

                        报文输入（HEX）

                    </span>




                    <Button


                        type="text"


                        size="small"


                        icon={<ClearOutlined />}


                        onClick={() => {


                            clear();


                        }}


                    >

                        清空


                    </Button>



                </div>







                <Input.TextArea



                    autoSize={{

                        minRows: 8,

                        maxRows: 12

                    }}



                    value={hex}



                    onChange={(e) => {


                        setHex(

                            e.target.value

                        )


                    }}



                    placeholder={
                        `
请输入HEX报文，例如：

7E020000450100000000...
`
                    }



                />



            </div>








            {/* 字节统计 */}



            <div

                style={{

                    marginTop: 8,

                    color: "#999",

                    fontSize: 12

                }}

            >


                当前长度：

                {
                    hex
                        ?
                        Math.ceil(
                            hex.replace(/\s/g, "").length / 2
                        )
                        :
                        0
                }

                {" Bytes"}


            </div>









            {/* 示例 */}



            <div

                style={{

                    marginTop: 20

                }}

            >



                <div

                    style={{

                        fontWeight: 500,

                        marginBottom: 8

                    }}

                >

                    报文示例

                </div>





                <Space wrap>





                    <Tag


                        color="blue"


                        style={{

                            cursor: "pointer"

                        }}



                        onClick={() => {


                            setHex(

                                mockParseResult.rawHex

                            )


                        }}


                    >

                        位置上传 0200


                    </Tag>






                    <Tag


                        style={{

                            cursor: "pointer"

                        }}


                        onClick={() => {


                            setHex(

                                "7E0002000000007E"

                            )


                        }}



                    >

                        终端心跳 0002


                    </Tag>







                    <Tag


                        style={{

                            cursor: "pointer"

                        }}


                        onClick={() => {


                            setHex(

                                "7E0801000000007E"

                            )


                        }}



                    >

                        报警信息 0801


                    </Tag>




                </Space>





            </div>









            {/* 解析按钮 */}

            <div

                style={{

                    marginTop: "auto",

                    paddingTop: 24

                }}

            >

                <Button



                    type="primary"



                    block



                    size="large"



                    icon={<PlayCircleOutlined />}



                    onClick={() => {


                        parse();


                    }}



                >



                    解析报文



                </Button>
            </div>





        </Card>



    )


}