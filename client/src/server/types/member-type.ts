import { ITimeStamps, TEntityId } from "./common";

export interface IMemberTypeResource extends ITimeStamps {
    id: TEntityId,

    name : string,
    description : string,
    prefix : string
}

