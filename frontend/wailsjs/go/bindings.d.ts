export interface go {
  "main": {
    "App": {
		Auth(arg1:string):Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
