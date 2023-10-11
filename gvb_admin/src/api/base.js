import {Service} from "@/services";

export function baseListApi(url, params){
    return Service.get(url, {params})
}