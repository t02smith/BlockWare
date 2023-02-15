export namespace games {
	
	export class Game {
	    title: string;
	    version: string;
	    release: string;
	    dev: string;
	    rootHash: number[];
	    previousVersion: number[];
	    IPFSId: string;
	    // Go type: big.Int
	    price?: any;
	    uploader: number[];
	
	    static createFrom(source: any = {}) {
	        return new Game(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.version = source["version"];
	        this.release = source["release"];
	        this.dev = source["dev"];
	        this.rootHash = source["rootHash"];
	        this.previousVersion = source["previousVersion"];
	        this.IPFSId = source["IPFSId"];
	        this.price = this.convertValues(source["price"], null);
	        this.uploader = source["uploader"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

