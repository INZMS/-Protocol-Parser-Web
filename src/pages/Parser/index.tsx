import { Layout } from "antd";


import Header from "../../components/AppHeader";
import InputPanel from "./InputPanel";
import ResultPanel from "./ResultPanel";
import HistoryTable from "./HistoryTable";


const { Content } = Layout;



export default function Parser() {


    return (

        <Layout
            style={{
                height: "100vh",
                background: "#f5f7fa"
            }}
        >


            <Header />


            <Content

                style={{

                    padding: 24,
                    display: "flex",


                    flexDirection: "column",


                    gap: 20,


                    overflow: "hidden",
                    flex: 1,

                    minHeight: 0

                }}

            >


                <div

                    style={{
                        flex: 1,

                        display: "flex",

                        gap: 20,

                        minHeight: 0,

                        alignItems: "stretch"

                    }}

                >


                    {/* 左侧 */}

                    <div

                        style={{

                            width: 360,
                            height: "100%",
                            minHeight: 0,
                            overflow: "hidden"

                        }}

                    >

                        <InputPanel />



                    </div>



                    {/* 右侧 */}

                    <div style={{
                        flex: 1,
                        height: 0,
                        minWidth: 0,
                        minHeight: 0,
                        overflow: "hidden"

                    }}>

                        <ResultPanel />

                    </div>

                </div>



                {/* 历史区域 */}

                <div
                    style={{

                        flex: "0 0 auto"

                    }}>
                    <HistoryTable />
                </div>



            </Content>


        </Layout>


    )

}