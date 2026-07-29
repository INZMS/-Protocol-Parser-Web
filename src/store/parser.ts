import { create } from "zustand";

import {
    mockParseResult
} from "../mock/parser";


interface ParserStore {


    hex: string;


    result: any;


    setHex: (hex: string) => void;


    parse: () => void;


    clear: () => void;


}



export const useParserStore = create<ParserStore>((set) => ({


    hex: "",


    result: null,



    setHex: (hex) => set({

        hex

    }),



    parse: () => set({

        result: mockParseResult

    }),



    clear: () => set({

        hex: "",

        result: null

    })



}))