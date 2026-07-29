import {
    Card,
    Statistic,
    Row,
    Col,
    Tabs,
    Button
} from "antd";


import {
    CopyOutlined
} from "@ant-design/icons";


import ParseTable from "../../components/ParseTable";


import {
    useParserStore
} from "../../store/parser";




export default function ResultPanel() {


    const {

        result

    } = useParserStore();





    /**
     * 未解析状态
     */

    if (!result) {


        return (

            <Card

                title="解析结果"

                style={{

                    height: "100%",

                    overflow: "hidden",

                    display: "flex",

                    flexDirection: "column"

                }}


                styles={{

                    body: {

                        flex: 1,

                        minHeight: 0,

                        display: "flex",

                        alignItems: "center",

                        justifyContent: "center"

                    }

                }}

            >


                <div

                    style={{

                        color: "#999",

                        textAlign: "center"

                    }}

                >

                    <div>

                        暂无解析结果

                    </div>


                    <div

                        style={{

                            marginTop: 8

                        }}

                    >

                        请输入HEX报文后点击解析

                    </div>


                </div>


            </Card>

        )

    }









    return (


        <Card

            title="解析结果"


            style={{

                height: "100%",

                overflow: "hidden",

                display: "flex",

                flexDirection: "column"

            }}


            styles={{


                body: {

                    flex: 1,

                    minHeight: 0,

                    overflow: "hidden",

                    display: "flex",

                    flexDirection: "column"


                }


            }}


        >





            {/* 基础信息 */}


            <Row

                gutter={16}

            >



                <Col span={6}>


                    <Statistic

                        title="协议"

                        value={result.protocol}

                    />


                </Col>





                <Col span={6}>


                    <Statistic

                        title="消息ID"

                        value={result.messageId}

                    />


                </Col>





                <Col span={6}>


                    <Statistic

                        title="消息名称"

                        value={result.messageName}

                    />


                </Col>





                <Col span={6}>


                    <Statistic

                        title="报文长度"

                        value={`${result.length} Bytes`}

                    />


                </Col>




            </Row>











            {/* 操作按钮 */}



            <div

                style={{


                    marginTop: 20,


                    marginBottom: 10,


                    textAlign: "right"


                }}

            >


                <Button

                    icon={<CopyOutlined />}

                    onClick={() => {


                        navigator.clipboard.writeText(


                            JSON.stringify(

                                result.json,

                                null,

                                2

                            )


                        )


                    }}

                >

                    复制结果

                </Button>



            </div>









            {/* Tabs区域 */}


            <div

                style={{


                    flex: 1,


                    minHeight: 0,
                    height: 0,


                    overflow: "hidden",





                }}

            >


                <Tabs


                    className="result-tabs"


                    style={{





                        height: "100%",



                    }}


                    defaultActiveKey="table"



                    items={[



                        {


                            key: "table",


                            label: "表格视图",


                            children: (


                                <div

                                    style={{


                                        height: "100%",


                                        overflow: "auto"


                                    }}

                                >

                                    <ParseTable />

                                </div>


                            )


                        },







                        {


                            key: "json",


                            label: "JSON视图",


                            children: (

                                <div

                                    style={{

                                        height: "100%",

                                        overflow: "auto"

                                    }}

                                >
                                    <pre


                                        style={{



                                            height: "100%",

                                            overflow: "auto",

                                            margin: 0

                                        }}


                                    >


                                        {

                                            JSON.stringify(

                                                result.json,

                                                null,

                                                2

                                            )

                                        }


                                    </pre>
                                </div>


                            )


                        }



                    ]}



                />


            </div>





        </Card>



    )

}