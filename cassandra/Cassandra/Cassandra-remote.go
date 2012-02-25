/* Autogenerated by Thrift Compiler (0.8.0-dev)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 */
package main

import (
        "flag"
        "fmt"
        "http"
        "net"
        "os"
        "strconv"
        "thrift"
        "thriftlib/cassandra"
)

func Usage() {
  fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:\n")
  flag.PrintDefaults()
  fmt.Fprint(os.Stderr, "Functions:\n")
  fmt.Fprint(os.Stderr, "  login(auth_request *AuthenticationRequest) (authnx *AuthenticationException, authzx *AuthorizationException, err error)\n")
  fmt.Fprint(os.Stderr, "  set_keyspace(keyspace string) (ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  get(key string, column_path *ColumnPath, consistency_level ConsistencyLevel) (retval1210 *ColumnOrSuperColumn, ire *InvalidRequestException, nfe *NotFoundException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  get_slice(key string, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval1211 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  get_count(key string, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval1212 int32, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  multiget_slice(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval1213 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  multiget_count(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval1214 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  get_range_slices(column_parent *ColumnParent, predicate *SlicePredicate, range_a1 *KeyRange, consistency_level ConsistencyLevel) (retval1215 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  get_indexed_slices(column_parent *ColumnParent, index_clause *IndexClause, column_predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval1216 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  insert(key string, column_parent *ColumnParent, column *Column, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  add(key string, column_parent *ColumnParent, column *CounterColumn, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  remove(key string, column_path *ColumnPath, timestamp int64, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  remove_counter(key string, path *ColumnPath, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  batch_mutate(mutation_map thrift.TMap, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)\n")
  fmt.Fprint(os.Stderr, "  truncate(cfname string) (ire *InvalidRequestException, ue *UnavailableException, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_schema_versions() (retval1223 thrift.TMap, ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_keyspaces() (retval1224 thrift.TList, ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_cluster_name() (retval1225 string, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_version() (retval1226 string, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_ring(keyspace string) (retval1227 thrift.TList, ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_partitioner() (retval1228 string, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_snitch() (retval1229 string, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_keyspace(keyspace string) (retval1230 *KsDef, nfe *NotFoundException, ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  describe_splits(cfName string, start_token string, end_token string, keys_per_split int32) (retval1231 thrift.TList, ire *InvalidRequestException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_add_column_family(cf_def *CfDef) (retval1232 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_drop_column_family(column_family string) (retval1233 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_add_keyspace(ks_def *KsDef) (retval1234 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_drop_keyspace(keyspace string) (retval1235 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_update_keyspace(ks_def *KsDef) (retval1236 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  system_update_column_family(cf_def *CfDef) (retval1237 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "  execute_cql_query(query string, compression Compression) (retval1238 *CqlResult, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, sde *SchemaDisagreementException, err error)\n")
  fmt.Fprint(os.Stderr, "\n")
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var help bool
  var url http.URL
  var trans thrift.TTransport
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.BoolVar(&help, "help", false, "See usage string")
  flag.Parse()
  if help || flag.NArg() == 0 {
    flag.Usage()
  }
  
  if len(urlString) > 0 {
    url, err := http.ParseURL(urlString)
    if err != nil {
      fmt.Fprint(os.Stderr, "Error parsing URL: ", err.Error(), "\n")
      flag.Usage()
    }
    host = url.Host
    useHttp = len(url.Scheme) <= 0 || url.Scheme == "http"
  } else if useHttp {
    _, err := http.ParseURL(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprint(os.Stderr, "Error parsing URL: ", err.Error(), "\n")
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(url.Raw)
  } else {
    addr, err := net.ResolveTCPAddr("tcp", fmt.Sprint(host, ":", port))
    if err != nil {
      fmt.Fprint(os.Stderr, "Error resolving address", err.Error())
      os.Exit(1)
    }
    trans, err = thrift.NewTNonblockingSocketAddr(addr)
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprint(os.Stderr, "Error creating transport", err.Error())
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprint(os.Stderr, "Invalid protocol specified: ", protocol, "\n")
    Usage()
    os.Exit(1)
  }
  client := cassandra.NewCassandraClientFactory(trans, protocolFactory)
  if err = trans.Open(); err != nil {
    fmt.Fprint(os.Stderr, "Error opening socket to ", host, ":", port, " ", err.Error())
    os.Exit(1)
  }
  
  switch cmd {
  case "login":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "Login requires 1 args\n")
      flag.Usage()
    }
    arg1239 := flag.Arg(1)
    mbTrans1240 := thrift.NewTMemoryBufferLen(len(arg1239))
    defer mbTrans1240.Close()
    _, err1241 := mbTrans1240.WriteString(arg1239)
    if err1241 != nil {
      Usage()
      return
    }
    factory1242 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1243 := factory1242.GetProtocol(mbTrans1240)
    argvalue0 := cassandra.NewAuthenticationRequest()
    err1244 := argvalue0.Read(jsProt1243)
    if err1244 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Login(value0))
    fmt.Print("\n")
    break
  case "set_keyspace":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SetKeyspace requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SetKeyspace(value0))
    fmt.Print("\n")
    break
  case "get":
    if flag.NArg() - 1 != 3 {
      fmt.Fprint(os.Stderr, "Get requires 3 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1247 := flag.Arg(2)
    mbTrans1248 := thrift.NewTMemoryBufferLen(len(arg1247))
    defer mbTrans1248.Close()
    _, err1249 := mbTrans1248.WriteString(arg1247)
    if err1249 != nil {
      Usage()
      return
    }
    factory1250 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1251 := factory1250.GetProtocol(mbTrans1248)
    argvalue1 := cassandra.NewColumnPath()
    err1252 := argvalue1.Read(jsProt1251)
    if err1252 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := cassandra.ConsistencyLevel(tmp2)
    value2 := argvalue2
    fmt.Print(client.Get(value0, value1, value2))
    fmt.Print("\n")
    break
  case "get_slice":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "GetSlice requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1254 := flag.Arg(2)
    mbTrans1255 := thrift.NewTMemoryBufferLen(len(arg1254))
    defer mbTrans1255.Close()
    _, err1256 := mbTrans1255.WriteString(arg1254)
    if err1256 != nil {
      Usage()
      return
    }
    factory1257 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1258 := factory1257.GetProtocol(mbTrans1255)
    argvalue1 := cassandra.NewColumnParent()
    err1259 := argvalue1.Read(jsProt1258)
    if err1259 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1260 := flag.Arg(3)
    mbTrans1261 := thrift.NewTMemoryBufferLen(len(arg1260))
    defer mbTrans1261.Close()
    _, err1262 := mbTrans1261.WriteString(arg1260)
    if err1262 != nil {
      Usage()
      return
    }
    factory1263 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1264 := factory1263.GetProtocol(mbTrans1261)
    argvalue2 := cassandra.NewSlicePredicate()
    err1265 := argvalue2.Read(jsProt1264)
    if err1265 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.GetSlice(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "get_count":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "GetCount requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1267 := flag.Arg(2)
    mbTrans1268 := thrift.NewTMemoryBufferLen(len(arg1267))
    defer mbTrans1268.Close()
    _, err1269 := mbTrans1268.WriteString(arg1267)
    if err1269 != nil {
      Usage()
      return
    }
    factory1270 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1271 := factory1270.GetProtocol(mbTrans1268)
    argvalue1 := cassandra.NewColumnParent()
    err1272 := argvalue1.Read(jsProt1271)
    if err1272 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1273 := flag.Arg(3)
    mbTrans1274 := thrift.NewTMemoryBufferLen(len(arg1273))
    defer mbTrans1274.Close()
    _, err1275 := mbTrans1274.WriteString(arg1273)
    if err1275 != nil {
      Usage()
      return
    }
    factory1276 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1277 := factory1276.GetProtocol(mbTrans1274)
    argvalue2 := cassandra.NewSlicePredicate()
    err1278 := argvalue2.Read(jsProt1277)
    if err1278 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.GetCount(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "multiget_slice":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "MultigetSlice requires 4 args\n")
      flag.Usage()
    }
    arg1279 := flag.Arg(1)
    mbTrans1280 := thrift.NewTMemoryBufferLen(len(arg1279))
    defer mbTrans1280.Close()
    _, err1281 := mbTrans1280.WriteString(arg1279)
    if err1281 != nil { 
      Usage()
      return
    }
    factory1282 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1283 := factory1282.GetProtocol(mbTrans1280)
    containerStruct0 := cassandra.NewMultigetSliceArgs()
    err1284 := containerStruct0.ReadField1(jsProt1283)
    if err1284 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Keys
    value0 := argvalue0
    arg1285 := flag.Arg(2)
    mbTrans1286 := thrift.NewTMemoryBufferLen(len(arg1285))
    defer mbTrans1286.Close()
    _, err1287 := mbTrans1286.WriteString(arg1285)
    if err1287 != nil {
      Usage()
      return
    }
    factory1288 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1289 := factory1288.GetProtocol(mbTrans1286)
    argvalue1 := cassandra.NewColumnParent()
    err1290 := argvalue1.Read(jsProt1289)
    if err1290 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1291 := flag.Arg(3)
    mbTrans1292 := thrift.NewTMemoryBufferLen(len(arg1291))
    defer mbTrans1292.Close()
    _, err1293 := mbTrans1292.WriteString(arg1291)
    if err1293 != nil {
      Usage()
      return
    }
    factory1294 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1295 := factory1294.GetProtocol(mbTrans1292)
    argvalue2 := cassandra.NewSlicePredicate()
    err1296 := argvalue2.Read(jsProt1295)
    if err1296 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.MultigetSlice(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "multiget_count":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "MultigetCount requires 4 args\n")
      flag.Usage()
    }
    arg1297 := flag.Arg(1)
    mbTrans1298 := thrift.NewTMemoryBufferLen(len(arg1297))
    defer mbTrans1298.Close()
    _, err1299 := mbTrans1298.WriteString(arg1297)
    if err1299 != nil { 
      Usage()
      return
    }
    factory1300 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1301 := factory1300.GetProtocol(mbTrans1298)
    containerStruct0 := cassandra.NewMultigetCountArgs()
    err1302 := containerStruct0.ReadField1(jsProt1301)
    if err1302 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Keys
    value0 := argvalue0
    arg1303 := flag.Arg(2)
    mbTrans1304 := thrift.NewTMemoryBufferLen(len(arg1303))
    defer mbTrans1304.Close()
    _, err1305 := mbTrans1304.WriteString(arg1303)
    if err1305 != nil {
      Usage()
      return
    }
    factory1306 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1307 := factory1306.GetProtocol(mbTrans1304)
    argvalue1 := cassandra.NewColumnParent()
    err1308 := argvalue1.Read(jsProt1307)
    if err1308 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1309 := flag.Arg(3)
    mbTrans1310 := thrift.NewTMemoryBufferLen(len(arg1309))
    defer mbTrans1310.Close()
    _, err1311 := mbTrans1310.WriteString(arg1309)
    if err1311 != nil {
      Usage()
      return
    }
    factory1312 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1313 := factory1312.GetProtocol(mbTrans1310)
    argvalue2 := cassandra.NewSlicePredicate()
    err1314 := argvalue2.Read(jsProt1313)
    if err1314 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.MultigetCount(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "get_range_slices":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "GetRangeSlices requires 4 args\n")
      flag.Usage()
    }
    arg1315 := flag.Arg(1)
    mbTrans1316 := thrift.NewTMemoryBufferLen(len(arg1315))
    defer mbTrans1316.Close()
    _, err1317 := mbTrans1316.WriteString(arg1315)
    if err1317 != nil {
      Usage()
      return
    }
    factory1318 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1319 := factory1318.GetProtocol(mbTrans1316)
    argvalue0 := cassandra.NewColumnParent()
    err1320 := argvalue0.Read(jsProt1319)
    if err1320 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg1321 := flag.Arg(2)
    mbTrans1322 := thrift.NewTMemoryBufferLen(len(arg1321))
    defer mbTrans1322.Close()
    _, err1323 := mbTrans1322.WriteString(arg1321)
    if err1323 != nil {
      Usage()
      return
    }
    factory1324 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1325 := factory1324.GetProtocol(mbTrans1322)
    argvalue1 := cassandra.NewSlicePredicate()
    err1326 := argvalue1.Read(jsProt1325)
    if err1326 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1327 := flag.Arg(3)
    mbTrans1328 := thrift.NewTMemoryBufferLen(len(arg1327))
    defer mbTrans1328.Close()
    _, err1329 := mbTrans1328.WriteString(arg1327)
    if err1329 != nil {
      Usage()
      return
    }
    factory1330 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1331 := factory1330.GetProtocol(mbTrans1328)
    argvalue2 := cassandra.NewKeyRange()
    err1332 := argvalue2.Read(jsProt1331)
    if err1332 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.GetRangeSlices(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "get_indexed_slices":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "GetIndexedSlices requires 4 args\n")
      flag.Usage()
    }
    arg1333 := flag.Arg(1)
    mbTrans1334 := thrift.NewTMemoryBufferLen(len(arg1333))
    defer mbTrans1334.Close()
    _, err1335 := mbTrans1334.WriteString(arg1333)
    if err1335 != nil {
      Usage()
      return
    }
    factory1336 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1337 := factory1336.GetProtocol(mbTrans1334)
    argvalue0 := cassandra.NewColumnParent()
    err1338 := argvalue0.Read(jsProt1337)
    if err1338 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg1339 := flag.Arg(2)
    mbTrans1340 := thrift.NewTMemoryBufferLen(len(arg1339))
    defer mbTrans1340.Close()
    _, err1341 := mbTrans1340.WriteString(arg1339)
    if err1341 != nil {
      Usage()
      return
    }
    factory1342 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1343 := factory1342.GetProtocol(mbTrans1340)
    argvalue1 := cassandra.NewIndexClause()
    err1344 := argvalue1.Read(jsProt1343)
    if err1344 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1345 := flag.Arg(3)
    mbTrans1346 := thrift.NewTMemoryBufferLen(len(arg1345))
    defer mbTrans1346.Close()
    _, err1347 := mbTrans1346.WriteString(arg1345)
    if err1347 != nil {
      Usage()
      return
    }
    factory1348 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1349 := factory1348.GetProtocol(mbTrans1346)
    argvalue2 := cassandra.NewSlicePredicate()
    err1350 := argvalue2.Read(jsProt1349)
    if err1350 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.GetIndexedSlices(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "insert":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "Insert requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1352 := flag.Arg(2)
    mbTrans1353 := thrift.NewTMemoryBufferLen(len(arg1352))
    defer mbTrans1353.Close()
    _, err1354 := mbTrans1353.WriteString(arg1352)
    if err1354 != nil {
      Usage()
      return
    }
    factory1355 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1356 := factory1355.GetProtocol(mbTrans1353)
    argvalue1 := cassandra.NewColumnParent()
    err1357 := argvalue1.Read(jsProt1356)
    if err1357 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1358 := flag.Arg(3)
    mbTrans1359 := thrift.NewTMemoryBufferLen(len(arg1358))
    defer mbTrans1359.Close()
    _, err1360 := mbTrans1359.WriteString(arg1358)
    if err1360 != nil {
      Usage()
      return
    }
    factory1361 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1362 := factory1361.GetProtocol(mbTrans1359)
    argvalue2 := cassandra.NewColumn()
    err1363 := argvalue2.Read(jsProt1362)
    if err1363 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.Insert(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "add":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "Add requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1365 := flag.Arg(2)
    mbTrans1366 := thrift.NewTMemoryBufferLen(len(arg1365))
    defer mbTrans1366.Close()
    _, err1367 := mbTrans1366.WriteString(arg1365)
    if err1367 != nil {
      Usage()
      return
    }
    factory1368 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1369 := factory1368.GetProtocol(mbTrans1366)
    argvalue1 := cassandra.NewColumnParent()
    err1370 := argvalue1.Read(jsProt1369)
    if err1370 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg1371 := flag.Arg(3)
    mbTrans1372 := thrift.NewTMemoryBufferLen(len(arg1371))
    defer mbTrans1372.Close()
    _, err1373 := mbTrans1372.WriteString(arg1371)
    if err1373 != nil {
      Usage()
      return
    }
    factory1374 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1375 := factory1374.GetProtocol(mbTrans1372)
    argvalue2 := cassandra.NewCounterColumn()
    err1376 := argvalue2.Read(jsProt1375)
    if err1376 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.Add(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "remove":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "Remove requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1378 := flag.Arg(2)
    mbTrans1379 := thrift.NewTMemoryBufferLen(len(arg1378))
    defer mbTrans1379.Close()
    _, err1380 := mbTrans1379.WriteString(arg1378)
    if err1380 != nil {
      Usage()
      return
    }
    factory1381 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1382 := factory1381.GetProtocol(mbTrans1379)
    argvalue1 := cassandra.NewColumnPath()
    err1383 := argvalue1.Read(jsProt1382)
    if err1383 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2, err1384 := (strconv.Atoi64(flag.Arg(3)))
    if err1384 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := cassandra.ConsistencyLevel(tmp3)
    value3 := argvalue3
    fmt.Print(client.Remove(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "remove_counter":
    if flag.NArg() - 1 != 3 {
      fmt.Fprint(os.Stderr, "RemoveCounter requires 3 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg1386 := flag.Arg(2)
    mbTrans1387 := thrift.NewTMemoryBufferLen(len(arg1386))
    defer mbTrans1387.Close()
    _, err1388 := mbTrans1387.WriteString(arg1386)
    if err1388 != nil {
      Usage()
      return
    }
    factory1389 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1390 := factory1389.GetProtocol(mbTrans1387)
    argvalue1 := cassandra.NewColumnPath()
    err1391 := argvalue1.Read(jsProt1390)
    if err1391 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := cassandra.ConsistencyLevel(tmp2)
    value2 := argvalue2
    fmt.Print(client.RemoveCounter(value0, value1, value2))
    fmt.Print("\n")
    break
  case "batch_mutate":
    if flag.NArg() - 1 != 2 {
      fmt.Fprint(os.Stderr, "BatchMutate requires 2 args\n")
      flag.Usage()
    }
    arg1392 := flag.Arg(1)
    mbTrans1393 := thrift.NewTMemoryBufferLen(len(arg1392))
    defer mbTrans1393.Close()
    _, err1394 := mbTrans1393.WriteString(arg1392)
    if err1394 != nil { 
      Usage()
      return
    }
    factory1395 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1396 := factory1395.GetProtocol(mbTrans1393)
    containerStruct0 := cassandra.NewBatchMutateArgs()
    err1397 := containerStruct0.ReadField1(jsProt1396)
    if err1397 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.MutationMap
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := cassandra.ConsistencyLevel(tmp1)
    value1 := argvalue1
    fmt.Print(client.BatchMutate(value0, value1))
    fmt.Print("\n")
    break
  case "truncate":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "Truncate requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Truncate(value0))
    fmt.Print("\n")
    break
  case "describe_schema_versions":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribeSchemaVersions requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribeSchemaVersions())
    fmt.Print("\n")
    break
  case "describe_keyspaces":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribeKeyspaces requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribeKeyspaces())
    fmt.Print("\n")
    break
  case "describe_cluster_name":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribeClusterName requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribeClusterName())
    fmt.Print("\n")
    break
  case "describe_version":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribeVersion requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribeVersion())
    fmt.Print("\n")
    break
  case "describe_ring":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "DescribeRing requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.DescribeRing(value0))
    fmt.Print("\n")
    break
  case "describe_partitioner":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribePartitioner requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribePartitioner())
    fmt.Print("\n")
    break
  case "describe_snitch":
    if flag.NArg() - 1 != 0 {
      fmt.Fprint(os.Stderr, "DescribeSnitch requires 0 args\n")
      flag.Usage()
    }
    fmt.Print(client.DescribeSnitch())
    fmt.Print("\n")
    break
  case "describe_keyspace":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "DescribeKeyspace requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.DescribeKeyspace(value0))
    fmt.Print("\n")
    break
  case "describe_splits":
    if flag.NArg() - 1 != 4 {
      fmt.Fprint(os.Stderr, "DescribeSplits requires 4 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err1404 := (strconv.Atoi(flag.Arg(4)))
    if err1404 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    fmt.Print(client.DescribeSplits(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "system_add_column_family":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemAddColumnFamily requires 1 args\n")
      flag.Usage()
    }
    arg1405 := flag.Arg(1)
    mbTrans1406 := thrift.NewTMemoryBufferLen(len(arg1405))
    defer mbTrans1406.Close()
    _, err1407 := mbTrans1406.WriteString(arg1405)
    if err1407 != nil {
      Usage()
      return
    }
    factory1408 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1409 := factory1408.GetProtocol(mbTrans1406)
    argvalue0 := cassandra.NewCfDef()
    err1410 := argvalue0.Read(jsProt1409)
    if err1410 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SystemAddColumnFamily(value0))
    fmt.Print("\n")
    break
  case "system_drop_column_family":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemDropColumnFamily requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SystemDropColumnFamily(value0))
    fmt.Print("\n")
    break
  case "system_add_keyspace":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemAddKeyspace requires 1 args\n")
      flag.Usage()
    }
    arg1412 := flag.Arg(1)
    mbTrans1413 := thrift.NewTMemoryBufferLen(len(arg1412))
    defer mbTrans1413.Close()
    _, err1414 := mbTrans1413.WriteString(arg1412)
    if err1414 != nil {
      Usage()
      return
    }
    factory1415 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1416 := factory1415.GetProtocol(mbTrans1413)
    argvalue0 := cassandra.NewKsDef()
    err1417 := argvalue0.Read(jsProt1416)
    if err1417 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SystemAddKeyspace(value0))
    fmt.Print("\n")
    break
  case "system_drop_keyspace":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemDropKeyspace requires 1 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SystemDropKeyspace(value0))
    fmt.Print("\n")
    break
  case "system_update_keyspace":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemUpdateKeyspace requires 1 args\n")
      flag.Usage()
    }
    arg1419 := flag.Arg(1)
    mbTrans1420 := thrift.NewTMemoryBufferLen(len(arg1419))
    defer mbTrans1420.Close()
    _, err1421 := mbTrans1420.WriteString(arg1419)
    if err1421 != nil {
      Usage()
      return
    }
    factory1422 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1423 := factory1422.GetProtocol(mbTrans1420)
    argvalue0 := cassandra.NewKsDef()
    err1424 := argvalue0.Read(jsProt1423)
    if err1424 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SystemUpdateKeyspace(value0))
    fmt.Print("\n")
    break
  case "system_update_column_family":
    if flag.NArg() - 1 != 1 {
      fmt.Fprint(os.Stderr, "SystemUpdateColumnFamily requires 1 args\n")
      flag.Usage()
    }
    arg1425 := flag.Arg(1)
    mbTrans1426 := thrift.NewTMemoryBufferLen(len(arg1425))
    defer mbTrans1426.Close()
    _, err1427 := mbTrans1426.WriteString(arg1425)
    if err1427 != nil {
      Usage()
      return
    }
    factory1428 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt1429 := factory1428.GetProtocol(mbTrans1426)
    argvalue0 := cassandra.NewCfDef()
    err1430 := argvalue0.Read(jsProt1429)
    if err1430 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.SystemUpdateColumnFamily(value0))
    fmt.Print("\n")
    break
  case "execute_cql_query":
    if flag.NArg() - 1 != 2 {
      fmt.Fprint(os.Stderr, "ExecuteCqlQuery requires 2 args\n")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := cassandra.Compression(tmp1)
    value1 := argvalue1
    fmt.Print(client.ExecuteCqlQuery(value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprint(os.Stderr, "Invalid function ", cmd, "\n")
  }
}