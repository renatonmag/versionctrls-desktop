export namespace backend {
	
	export class OpenRepositoryResult {
	    Path: string;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new OpenRepositoryResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Error = source["Error"];
	    }
	}

}

export namespace git {
	
	export class RepositoryInfo {
	    Path: string;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new RepositoryInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Name = source["Name"];
	    }
	}

}

