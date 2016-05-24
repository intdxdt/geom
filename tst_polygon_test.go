package geom

import (
    . "github.com/franela/goblin"
    . "simplex/util/math"
    "testing"
    "fmt"
)

func TestPolygon(t *testing.T) {
    g := Goblin(t)

    sh := []*Point{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}, }
    h1 := []*Point{{20, 30}, {35, 35}, {30, 20}, {20, 30}, }
    poly0 := NewPolygon(sh)
    poly := NewPolygon(sh, h1)
    //poly := NewPolygon(sh)

    g.Describe("Polygon", func() {
        g.It("should test polygon relates", func() {
            wkt := "POLYGON (( 33.52991674117594 27.137460594059416, 33.52991674117594 30.589750223527805, 36.44941148514852 30.589750223527805, 36.44941148514852 27.137460594059416, 33.52991674117594 27.137460594059416 ))"
            ply_inpoly := NewPolygonFromWKT(wkt)
            ply_inpoly_clone := ply_inpoly.Clone()
            g.Assert(poly.Envelope().Equals(poly.Shell.BBox())).IsTrue()
            g.Assert(poly.Envelope().Equals(poly.BBox())).IsTrue()
            g.Assert(poly.Intersects(ply_inpoly)).IsTrue()
            g.Assert(poly.Intersects(ply_inpoly_clone)).IsTrue()
            //g.Assert(poly.Contains(ply_inpoly)).IsTrue()
        })
        g.It("should test polygon string", func() {
            g.Assert(poly.String()).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))",
            )
            g.Assert(fmt.Sprintf("%v", poly0)).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10))",
            )
            g.Assert(fmt.Sprintf("%v", poly)).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))",
            )
        })

        g.It("should test area polygon with holes/chull", func() {
            ply := "POLYGON (( 5.6 7.9, 5.6 8.9, 6.6 8.9, 6.6 7.9, 5.6 7.9 ), (5.81 8.43, 5.84 8.68, 6.30 8.77, 6.42 8.52, 6.36 8.19, 5.96 8.14, 5.81 8.43 ))"
            gply := NewPolygonFromWKT(ply)
            g.Assert(Round(gply.Area(), 5)).Equal(0.7025)
            arr  := []*Point{{5.89, 8.64}, {5.87, 8.74}, {5.91, 8.86},
                {6.1310016380016386, 8.896833606333606}, {6.33, 8.93},
                {6.36, 8.77}, {6.58, 8.78}, {6.65, 8.61}, {6.57, 8.36},
                {6.32, 8.5 }, {6.21, 8.34}, {5.94, 8.29}, {5.78, 8.5 },
                {5.89, 8.64}}
            harr := ConvexHull(arr)
            sharr := SimpleHull(arr)

            hply := NewPolygon(harr)
            sply := NewPolygon(sharr)
            g.Assert(len(harr)).Equal(7)

            g.Assert(Round(hply.Area(), 6)).Equal(0.41995)
            g.Assert(Round(sply.Area(), 6)).Equal(0.41995)

            arr2 := []*Point {
                {5.89, 8.64 }, {5.78, 8.5  }, {5.94, 8.29 }, {6.21, 8.34 },
                {6.32, 8.5  }, {6.57, 8.36 }, {6.65, 8.61 }, {6.58, 8.78 },
                {6.36, 8.77 }, {6.33, 8.93 }, {6.1310016380016386, 8.896833606333606},
                {5.91, 8.86 }, {5.87, 8.74 }, {5.89, 8.64 },
            }
            harr = ConvexHull(arr2)
            sharr = ConvexHull(arr2)

            hply = NewPolygon(harr)
            sply = NewPolygon(sharr)

            g.Assert(len(harr)).Equal(7)
            g.Assert(len(sharr)).Equal(7)
            g.Assert(Round(hply.Area(), 6)).Equal(0.41995)
            g.Assert(Round(sply.Area(), 6)).Equal(0.41995)


            arr3 := []*Point {
                {5.01, 8.74} , {5.01, 8.74}, {5.10, 8.59}, {5.29, 8.62}, {5.30, 8.50},
                {5.65, 8.58}, {5.46, 8.66}, {5.67, 8.79}, {5.67, 8.79}, {5.63, 9.07},
                {5.27, 9.04}, {5.40, 8.88}, {5.15, 8.84}, {5.15, 8.84}, {5.02, 8.76},
                {5.01, 8.74},
            }

            harr = ConvexHull(arr3)
            sharr = ConvexHull(arr3)

            hply = NewPolygon(harr)
            sply = NewPolygon(sharr)

            g.Assert(len(harr)).Equal(8)
            g.Assert(len(sharr)).Equal(8)
            g.Assert(Round(hply.Area(), 6)).Equal(0.27795)
            g.Assert(Round(sply.Area(), 6)).Equal(0.27795)


            wkt_data := "POLYGON (( 2958200.37460972 4491861.80361633, 2958440.524147205 4491771.742896985, 2958650.6508180276 4491855.793940935, 2958392.489786923 4491435.5314368345, 2958344.4665586017 4490973.249489654, 2958254.409090545 4490570.993055433, 2958440.524147205 4490372.868021447, 2958896.81160802 4490300.818387367, 2959184.9843738303 4490192.758868743, 2959377.1106829904 4490192.758868743, 2959575.2371127047 4490342.8519493975, 2959755.3409168646 4490282.806056019, 2959821.3867707513 4490054.67120602, 2959911.4442388006 4490156.734607343, 2960109.570668515 4490162.743311103, 2960169.6052699015 4489982.623753354, 2960373.73182017 4489904.582240969, 2960397.7434343323 4490018.647417948, 2960613.881357655 4490120.710469432, 2960739.9618129283 4490048.662564043, 2960727.950439874 4489856.541899648, 2960625.8927307166 4489856.541899648, 2960499.812275443 4489766.484358743, 2960649.904344879 4489478.305413313, 2960757.9733065367 4489184.126221187, 2960884.05376181 4489154.11354389, 2960938.0882426426 4489256.168200422, 2961064.168697916 4489178.1180768795, 2961040.1459518 4488998.015391916, 2960962.099856805 4488757.864606198, 2961052.1573248617 4488547.730167106, 2961088.1803120784 4488229.532583475, 2960974.1112298667 4488091.44270239, 2961016.1343376376 4487857.292005297, 2961052.1573248617 4487617.139270764, 2961172.2265276276 4487623.146523006, 2961322.329729017 4487743.222273178, 2961502.4335331656 4487479.057432961, 2961754.594443716 4487358.9846996665, 2961814.6290450953 4487214.899229713, 2962054.778582584 4487106.829421867, 2962078.8013286963 4486902.692814201, 2962306.939493127 4486866.679825705, 2962667.169365339 4486824.646159198, 2962709.19247311 4486662.547880143, 2962661.158112839 4486422.394428901, 2962469.0429356247 4486392.375633124, 2962427.0198278464 4486170.2336072065, 2962126.8245570287 4486182.246461991, 2961892.6862720437 4486254.295876682, 2961814.6290450953 4486476.425676972, 2961670.548228167 4486416.387862626, 2961460.410425391 4486152.228351723, 2961442.3989317827 4485966.115180336, 2961340.341222625 4485972.121489432, 2961130.2034198567 4486206.258211695, 2961124.2032993026 4486386.369083997, 2961232.27226096 4486440.414148316, 2961238.2723815143 4486602.5090031475, 2961112.191926241 4486824.646159198, 2960968.099977359 4486998.760724742, 2960769.9846795984 4486998.760724742, 2960493.8010229394 4487154.857198816, 2960277.674231559 4487160.864186946, 2960079.547801852 4487256.920453876, 2959971.478840187 4487533.094111022, 2959755.3409168646 4487533.094111022, 2959581.237233259 4487389.00625509, 2959455.1567779854 4487389.00625509, 2959221.0073610544 4487623.146523006, 2959088.926785227 4487665.169378828, 2958962.8463299535 4487623.146523006, 2958878.8001144044 4487707.206405647, 2958746.7084066346 4487719.207013436, 2958686.6738052443 4487653.168832723, 2958554.5932294168 4487929.338895369, 2958452.5243883096 4487941.353760056, 2958362.47805221 4488019.394701414, 2958242.397717491 4488103.44374878, 2958110.3171416633 4488091.44270239, 2957948.213699173 4487893.322390255, 2957676.916968163 4487873.605786528, 2957654.7309936434 4488101.427235939, 2957565.4416300803 4488149.935679749, 2957468.827444017 4488147.120951578, 2957428.796955131 4488137.472461212, 2957280.307886362 4488045.791239213, 2957228.766962122 4488005.853373263, 2957176.524725098 4487936.312556416, 2957242.5037872866 4487749.6636932865, 2957299.811061144 4487491.071783088, 2957524.665300604 4487345.038148042, 2957631.7657826915 4487163.944695003, 2957434.652360346 4487012.04871792, 2957285.461978782 4487023.124392081, 2957169.522729125 4487067.455209766, 2957137.707618654 4487034.774165612, 2957028.3250869997 4486921.021400422, 2956921.5251675323 4486988.0071240105, 2956710.441149097 4486946.883103915, 2956580.3086643554 4486964.399648484, 2956399.9042975716 4487044.197609656, 2956369.402757097 4487057.177612528, 2956336.8640699387 4487047.068051189, 2956246.4615114667 4486970.47653807, 2956176.6975865886 4486865.71369357, 2956166.489589285 4486806.471761156, 2956164.029428534 4486768.512853928, 2956183.321096286 4486720.108795125, 2956286.1023821346 4486635.776771374, 2956442.205704078 4486528.398987088, 2956696.1365945227 4486312.960650049, 2956734.0742769875 4486214.574830495, 2956613.2035738863 4486229.2059350945, 2956582.368074935 4486194.287332542, 2956630.9924285114 4486162.743081156, 2956648.692227546 4486079.269579906, 2956585.985958386 4486078.415527049, 2956581.8448733278 4486024.610336591, 2956603.140291918 4485992.632570837, 2956682.8784431703 4485744.752181314, 2956646.1318792626 4485637.76150199, 2956652.7219931185 4485496.961634614, 2956708.459662158 4485356.471650045, 2956652.7219931185 4485326.245898671, 2956515.0643107966 4485401.537440658, 2956398.5016719922 4485502.365671504, 2956245.3037887663 4485569.594362542, 2956159.8772115223 4485513.859759193, 2956112.9003864154 4485485.957568087, 2956068.606361024 4485384.849495683, 2956056.0383905135 4485291.288294539, 2956093.508531116 4485180.158324789, 2956176.486079555 4485121.401872754, 2956361.0649272352 4485119.20393702, 2956494.971142713 4484922.021572173, 2956700.055040598 4484854.783020917, 2956821.994410813 4484758.174728554, 2956533.0758044124 4484481.807956379, 2956451.623332996 4484690.0692855045, 2956254.6991537847 4484771.487889297, 2956073.727057606 4484678.954071082, 2956095.0670039915 4484527.97590499, 2955979.717747625 4484458.066118699, 2955967.505999487 4484300.973776508, 2955774.63384974 4484221.210624784, 2955647.1507688835 4484318.359928127, 2955608.122155413 4484481.625972748, 2955487.997292895 4484580.345554691, 2955264.1115330085 4484702.93439091, 2955197.553609468 4484804.385887679, 2955157.7234956585 4484932.941052664, 2955098.6573738456 4484948.648325261, 2955128.3685459346 4485025.981086012, 2955184.339985907 4485006.955830194, 2955217.05678425 4485084.2750201, 2955163.857199602 4485107.192296393, 2955178.484580692 4485141.939294368, 2955096.1860811487 4485178.520363491, 2955049.353971377 4485073.901364107, 2955058.137079194 4485048.66024388, 2955050.823388651 4485042.444469821, 2955042.4076351486 4485072.809400916, 2955093.9819552302 4485199.001897864, 2955053.7844871096 4485213.771586835, 2954930.25324817 4485195.473975223, 2954829.520240959 4485124.383786283, 2954861.112712443 4485114.682070699, 2954834.6298055835 4484903.234492511, 2954860.589510832 4484786.928937003, 2954894.9761015438 4484782.883184321, 2954809.0151907466 4484661.819305111, 2954730.6685331278 4484535.4932645, 2954723.800120551 4484501.798177294, 2954726.727823153 4484462.293729741, 2954745.3181781173 4484435.220250014, 2954781.619464062 4484387.736838542, 2954804.929765437 4484357.457876455, 2954861.3130875267 4484398.025826082, 2954936.3869521134 4484336.264054708, 2955058.2706625834 4484261.526076917, 2955027.457427539 4484203.376685869, 2955138.3093764633 4484165.091212697, 2955172.1393697187 4484204.7065311335, 2955182.3807628714 4484199.807102043, 2955156.1204949953 4484158.414016474, 2955240.211238336 4484014.092425257, 2955424.389335852 4483826.309342172, 2955746.837372888 4483528.424917955, 2955919.0931529365 4483527.501084827, 2956217.2624090314 4483465.142536212, 2956406.694786515 4483213.570488296, 2956402.2420068793 4483165.364746906, 2956348.207526047 4483051.289902974, 2956393.625878293 4482861.676454034, 2956424.7062801234 4482732.628729343, 2956604.7544245347 4482637.9854062535, 2956707.836273007 4482632.400888935, 2956796.6692266613 4482659.567704022, 2956799.652589012 4482720.549831342, 2956772.9470431767 4482768.109697629, 2956765.143546868 4482857.197547704, 2956931.7888245843 4482936.026577592, 2957106.026091572 4482683.22150534, 2957083.873512905 4482629.475666769, 2957112.749788817 4482574.218512684, 2957180.565622613 4482563.819339078, 2957317.2102975585 4482503.733871017, 2957366.268797148 4482399.75728998, 2957472.779285945 4482321.7580335885, 2957589.453244243 4482419.967491277, 2957861.1618573703 4482111.696161047, 2958047.2769140303 4482045.665185388, 2957957.2194459736 4481961.608548742, 2957987.2423126437 4481877.552583486, 2958197.368983466 4481565.364585221, 2958197.368983466 4481421.27448111, 2958041.276793476 4481421.27448111, 2957975.2309395894 4481499.32304262, 2957819.138749592 4481499.32304262, 2957612.307135705 4481193.3177774325, 2957393.486412648 4481258.545721561, 2957317.6333116256 4481212.056250781, 2957314.7390048616 4481151.600653738, 2957345.374128733 4481107.700591765, 2957298.263720222 4481082.357030809, 2957259.246238701 4481132.134585172, 2957290.304376632 4481182.654075604, 2957212.3362052813 4481282.980065163, 2957012.3507400714 4481467.233076271, 2956926.990954537 4481482.277422495, 2956867.301443573 4481461.383280478, 2956767.904270243 4481414.487068556, 2956756.438362688 4481403.0114531815, 2956735.9667083286 4481376.827504445, 2956699.0086373873 4481299.423598666, 2956694.7451008894 4481277.242327608, 2956692.963989038 4481267.991981689, 2956730.7792200632 4481190.8967523165, 2956724.1891062073 4481101.948955517, 2956685.0491732433 4481047.245594136, 2956676.22153762 4480896.1935830675, 2956706.1998764947 4480790.665573843, 2956672.091584511 4480747.816550277, 2956664.7444981225 4480738.020912949, 2956661.2935939096 4480721.158444341, 2956663.575643465 4480675.552948788, 2956684.9935134985 4480630.983178575, 2956727.562086772 4480576.786032528, 2956815.582408149 4480544.307060655, 2956871.7987509966 4480544.894787826, 2956882.340706773 4480611.070277877, 2956914.2114769854 4480631.878770217, 2956993.715857312 4480589.226302933, 2956929.195080448 4480534.357684325, 2956896.400358461 4480457.05792756, 2956855.991383303 4480400.258920658, 2956696.426025197 4480232.522425398, 2956646.3879140913 4480147.178921051, 2956659.2230513766 4480101.212096449, 2956605.333285887 4479976.228125811, 2956674.128731195 4479861.880006954, 2956681.7318524197 4479715.784321655, 2956656.4734599553 4479444.377194364, 2956691.1049535424 4479326.886677541, 2956762.3271637484 4479243.649280667, 2956850.870686725 4479206.081432559, 2957008.7105927244 4479145.511430152, 2957152.802541606 4479055.251719654, 2957317.611047726 4479030.822546471, 2957412.1101634577 4479000.60096737, 2957735.081402097 4478878.680162925, 2957867.1619779244 4478710.576159038, 2958071.288528193 4478602.509717282, 2958065.2884076387 4478272.301136702, 2957987.2423126437 4478230.273806676, 2957903.1849651486 4477936.087193102, 2958005.253806252 4477804.008079294, 2958005.253806252 4477503.817902811, 2957939.2079523653 4477263.663528938, 2957975.2309395894 4477071.5523660295, 2958095.3112743087 4476933.4549605325, 2958095.3112743087 4476831.3938601315, 2957939.2079523653 4476843.410048515, 2957855.1506048664 4477005.511028755, 2957689.295695532 4476930.531321503, 2957452.9978124276 4476939.148365188, 2957339.507591564 4476934.112429712, 2957085.031235613 4476972.161779616, 2956782.0418455712 4476922.110126529, 2956530.537720021 4476649.418126509, 2956451.044471644 4476663.658254251, 2956359.8404128402 4476709.721982025, 2956133.316381026 4476876.395185597, 2956044.049281355 4477146.267286472, 2955844.2753231786 4477335.288189974, 2955204.978619501 4477370.639056716, 2954753.188466113 4477502.600819591, 2954562.787609067 4477684.171604738, 2954768.0162222907 4477947.447109513, 2954927.5927123465 4477917.956120376, 2954967.990555551 4477953.616724424, 2954994.7517611384 4478004.722539749, 2955013.3532480486 4478088.048074648, 2955003.5682648085 4478162.770207245, 2954915.080401581 4478166.925338589, 2954850.9937707298 4478285.242342338, 2954771.0107165948 4478324.038073573, 2955009.3791422285 4478665.539012842, 2954985.3563961163 4478911.69965842, 2955165.471332215 4478809.633374605, 2955375.609134987 4478821.637829304, 2955432.7939574122 4478858.182943452, 2955557.4829190485 4478834.565718893, 2955823.5587659404 4479073.958404336, 2955948.114144195 4479002.237967346, 2956107.100640945 4479082.311362527, 2956233.3146796003 4479196.776939683, 2956272.432348665 4479325.025756951, 2956285.4678610414 4479456.396364875, 2956233.7154297717 4479575.833244909, 2956176.6975865886 4479578.827579599, 2956127.3941841163 4479691.605478384, 2956071.6119872816 4479747.295258611, 2955967.2054368593 4479773.349221762, 2955834.5237357877 4479727.495968826, 2955652.5831600316 4479678.382696964, 2955453.2322159223 4479676.689622443, 2955324.9142388813 4479872.668304313, 2955170.04656329 4480036.760929413, 2954981.115123518 4480118.199508395, 2954741.80048221 4480192.60021795, 2954596.5062828287 4480080.866418384, 2954437.018848367 4480035.51556772, 2954378.4981920607 4479922.845926028, 2954286.503764864 4479865.699986465, 2954336.9537580907 4479542.098065689, 2953826.631816447 4479194.832091831, 2953430.390088968 4479253.91924832, 2953105.3037800044 4479242.082202014, 2953004.125494823 4479392.005113315, 2953032.3672496416 4479501.506854795, 2953042.5975108445 4479593.001708198, 2953004.125494823 4479722.206836231, 2952902.535327524 4479915.975523651, 2952808.692996789 4480234.243575495, 2952985.034202151 4480108.95018103, 2953398.218756132 4480349.53328665, 2953458.1643019244 4480542.865729891, 2953480.539519571 4480763.55955793, 2953431.9262979403 4481069.986148566, 2953315.5862981156 4481326.125152901, 2953318.0019310676 4481559.542724591, 2953311.734643735 4481627.991775248, 2953384.7713616453 4481676.470374454, 2953443.314281851 4481795.456612829, 2953403.317188807 4481825.5462652, 2953539.015648082 4481942.784768578, 2953575.528441068 4481927.039996903, 2953636.5983137153 4482012.03405289, 2953675.0035380423 4482107.105609275, 2953688.0056545623 4482249.819522634, 2953652.7396398783 4482451.948394656, 2953508.157885235 4482719.850011487, 2953380.3074500635 4482980.620107818, 2953274.876760334 4483036.999141604, 2953247.2806585617 4483147.056700904, 2953548.199506074 4483213.738452937, 2953564.57460317 4483386.701465927, 2953653.0402025096 4483473.596961997, 2953547.0974431187 4483555.160123937, 2953595.0984075516 4483624.657979831, 2953546.4740539677 4483650.007664494, 2953512.811039958 4483759.721692011, 2953675.905225914 4483776.099107347, 2953815.076853305 4483709.3158070035, 2953943.2501149997 4483711.569439027, 2954029.3223452866 4483773.397531945, 2954140.341273453 4483788.123225957, 2954209.147850711 4483885.89843753, 2954218.6545352265 4483982.988652036, 2954105.8544952087 4484176.87780422, 2953776.2597468607 4484392.566362049, 2953639.069606412 4484540.812814515, 2953563.5393319093 4484548.582162134, 2953467.3258960135 4484627.21388248, 2953430.390088968 4484765.3562737405, 2953411.5770950243 4484992.2984269895, 2953551.639278345 4485195.20798111, 2953645.7933036573 4485391.793469571, 2953624.1861904897 4485431.875408452, 2953667.578528002 4485443.775419716, 2953759.072017487 4485309.642072208, 2953860.9961432554 4485243.843017682, 2953940.1443012096 4485295.054246306, 2953914.641005866 4485325.685904093, 2953914.128936209 4485618.4271339215, 2954007.9156072065 4485634.933445938, 2954228.283671178 4485566.122316323, 2954305.88448821 4485587.8506243825, 2954334.983403109 4485638.083508424, 2954390.709940195 4485694.420785315, 2954572.3499533236 4485557.974212378, 2954629.7128869332 4485542.336030558, 2954697.1836303025 4485599.876814187, 2954747.6670193747 4485613.2050406, 2954812.3325115778 4485701.658971246, 2954876.998003777 4485661.3099982925, 2954960.4096982293 4485747.9022766985, 2955009.3791422285 4485966.115180336, 2954923.4850231335 4486170.317613136, 2954663.9324983954 4486437.41786027, 2954498.411547538 4486326.681813713, 2954227.949712705 4486171.689710088, 2954081.831749089 4486427.812942684, 2954060.2691637278 4486691.92336975, 2954108.81559366 4486908.6996582635, 2954234.8960489333 4487058.801864278, 2954042.769739773 4487196.87818303, 2954090.8041000515 4487437.035321005, 2953868.666056175 4487707.206405647, 2954012.746873107 4487809.261019535, 2953904.6890433915 4488055.425641928, 2953724.5741072856 4488109.451278895, 2953748.5857214555 4488349.6012558155, 2953820.6316958964 4488751.85670552, 2954132.8272078224 4488835.911609307, 2954288.919397816 4488739.8409144655, 2954457.03409281 4488739.8409144655, 2954511.0685736425 4488919.953180563, 2954354.965251699 4489298.197584726, 2954186.861688655 4489514.327349383, 2953790.60882923 4489616.3995090835, 2953646.5280122906 4489760.489887297, 2953622.5052661784 4489958.603376877, 2953826.631816447 4490120.710469432, 2953790.60882923 4490498.942062981, 2953634.5166392364 4490540.976417482, 2953568.4707853533 4490895.200620383, 2953772.5973356143 4491165.362555478, 2953772.5973356143 4491369.497917984, 2953604.49377257 4491519.579791255, 2953604.49377257 4491639.658040002, 2953748.5857214555 4491789.757780794, 2953748.5857214555 4492005.896769542, 2953976.7238858826 4491927.840229303, 2954192.861809209 4491879.81865456, 2954216.884555325 4491747.732438903, 2954433.011346698 4491765.74728154, 2954688.1778834946 4491507.574843474, 2954946.3389145955 4491393.507511742, 2955114.4424776398 4491153.358012389, 2955090.4308634773 4490745.099650577, 2954904.3046748675 4490727.086557712, 2954946.3389145955 4490456.921883319, 2955210.500066243 4490378.876848783, 2955288.546161242 4490438.909284413, 2955540.707071785 4490402.884179238, 2955678.7989001125 4490312.835963078, 2955912.937185101 4490420.8967163935, 2956083.868263215 4490575.559289709, 2956190.2896964103 4490509.517167989, 2956400.3273116387 4490530.737443931, 2956448.3950677663 4490628.267223574, 2956431.0626230463 4490721.525756899, 2956325.253447052 4490916.925938368, 2956412.950941898 4490906.196343079, 2956702.1923748218 4490779.248973589, 2956852.161992822 4490757.972167429, 2956877.832267396 4490799.825465035, 2956892.6488916203 4490852.128314231, 2956883.698804565 4490915.413148809, 2956734.4193674102 4491008.268014953, 2956623.244591955 4491048.063208763, 2956597.195831105 4491443.838185921, 2956711.008878492 4491450.954259504, 2957044.266038079 4491372.20144593, 2957299.811061144 4491387.498104189, 2957605.9953205734 4491621.643413559, 2957690.0526680723 4491867.799286589, 2957599.995200023 4491999.887011722, 2957672.041174464 4492077.930078376, 2957599.995200023 4491999.887011722, 2958200.37460972 4491861.80361633 ))"
            poly := NewPolygonFromWKT(wkt_data)
            ply_coords := poly.Shell.Coordinates()
            hull := ConvexHull(ply_coords)
            shull := ConvexHull(ply_coords)

            phull := NewPolygon(hull)
            pshull := NewPolygon(shull)
            g.Assert(Round(phull.Area(),  2)). Equal(110190309.57)
            g.Assert(Round(pshull.Area(), 2)).Equal(110190309.57)

            arr4 := []*Point{{5.10, 8.59},{5.10, 8.59},{5.10, 8.59},{5.10, 8.59}, {5.01, 8.74}, {5.01, 8.74}, {5.01, 8.74}, {5.01, 8.74}}
            hull = ConvexHull(arr4)
            phull = NewPolygon(hull)
            g.Assert(Round(phull.Area(), 2)).Equal(0.0)

        })

    })
}