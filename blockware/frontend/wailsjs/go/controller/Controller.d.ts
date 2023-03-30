// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {map[string]controller} from '../models';
import {controller} from '../models';

export function ConnectFromFile(arg1:string):Promise<void>;

export function ConnectToManyPeers(arg1:string):Promise<void>;

export function ConnectToPeer(arg1:string,arg2:number):Promise<void>;

export function ContinueAllDownloads():Promise<void>;

export function ContinueDownload(arg1:string):Promise<void>;

export function CreateDownload(arg1:string):Promise<void>;

export function DeployLibraryInstance(arg1:string):Promise<string>;

export function Disconnect(arg1:string,arg2:number):Promise<void>;

export function FetchOwnedGame(arg1:string):Promise<void>;

export function GetContractAddress():Promise<string>;

export function GetDirectory():Promise<string>;

export function GetDownloads():Promise<map[string]controller.ControllerDownload>;

export function GetGameFromStoreByRootHash(arg1:string):Promise<controller.ControllerGame>;

export function GetOwnedGames():Promise<Array<controller.ControllerGame>>;

export function GetPeerInformation():Promise<Array<controller.ControllerPeerData>>;

export function GetStoreGames():Promise<Array<controller.ControllerGame>>;

export function IsDownloading(arg1:string):Promise<number>;

export function JoinLibraryInstance(arg1:string,arg2:string):Promise<void>;

export function LoadDeferredRequests():Promise<void>;

export function PurchaseGame(arg1:string):Promise<void>;

export function ResendValidation(arg1:string,arg2:number):Promise<void>;

export function SelectFolder():Promise<string>;

export function SelectTxtFile():Promise<string>;

export function UninstallGame(arg1:string):Promise<void>;

export function UploadGame(arg1:string,arg2:string,arg3:string,arg4:string,arg5:number,arg6:number,arg7:number,arg8:string):Promise<void>;
