package jdapi

// PositionCreateParams 创建推广位参数
type PositionCreateParams struct {
	Code          string `json:"code"`          //渠道编码，联系京佣获取
	SubjectName   string `json:"subjectName"`   //主体名称，不能超过30字
	SubjectType   string `json:"subjectType"`   //主体类型（种类⻅附录7）
	ContactPerson string `json:"contactPerson"` //联系
	ContactPhone  string `json:"contactPhone"`  //联系⽅式（⼿机号）
	ChannelRatio  int    `json:"channelRatio"`  //给⼆级渠道设置的分佣⽐例
}

// PositionCreateResp 创建推广位返回数据
type PositionCreateResp struct {
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		Code        string `json:"code"`        //渠道编码
		PID         string `json:"pId"`         //pid
		PositionID  int64  `json:"positionId"`  //推广位id
		SubjectName string `json:"subjectName"` //主题名称
	} `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
}

// PositionQueryParams 查询下级渠道推广位参数
type PositionQueryParams struct {
	Code string `json:"code"` //渠道编码，联系京佣获取
}

// PositionQueryResp 查询下级渠道推广位返回数据
type PositionQueryResp struct {
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		Code        string `json:"code"`        //渠道编码
		PID         string `json:"pId"`         //pid
		PositionID  int64  `json:"positionId"`  //推广位id
		SubjectName string `json:"subjectName"` //主题名称
	} `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
}

// GoodsQueryParams 商品查询参数
type GoodsQueryParams struct {
	PageNo    int    `json:"pageNo"`              //分⻚参数，⻚码编号
	PageSize  int    `json:"pageSize"`            //分⻚参数，每⻚展示条数，最多60
	Code      string `json:"code"`                //渠道编码，联系来客商务获取
	Attribute int    `json:"attribute,omitempty"` //-1：全部商品，2：直播专区（默认），6：热销榜 单。默认为2直播专区
	//商品库⼆级编号  当attribute传6时，此时该参数可选：
	//600：热销爆品，601：今⽇必推，
	//602：秒杀商品，603：品牌好货，
	//604:9.9包邮
	SubAttribute int     `json:"subAttribute,omitempty"`
	Keywords     string  `json:"keywords,omitempty"`     //关键词，可以是落地 ⻚，也可以是sku，商 品名称
	ShopIdList   []int   `json:"shopIdList,omitempty"`   //店铺集合
	SkuIdList    []int   `json:"skuIdList,omitempty"`    //商品SKU集合，可只传 ⼀个
	Cat1Id       int     `json:"cat1Id,omitempty"`       //⼀级类⽬Id
	Cat2Id       int     `json:"cat2Id,omitempty"`       //⼆级类⽬Id
	Cat3Id       int     `json:"cat3Id,omitempty"`       //三级类⽬Id
	Owner        string  `json:"owner,omitempty"`        //g为⾃营，p为pop
	IsPinGou     int     `json:"isPinGou,omitempty"`     //是否拼购商品，1： 是，默认搜索全部
	HasCoupon    int     `json:"hasCoupon,omitempty"`    //是否有优惠券（默认有，如果允许每⻚优惠券则传0）
	DeliveryType int     `json:"deliveryType,omitempty"` //是否京配，1：是，默 认搜索全部
	IsCare       int     `json:"isCare,omitempty"`       //是否爱⼼咚咚商品，1：是，默认搜索全部
	FromPrice    float64 `json:"fromPrice,omitempty"`    //单价最⼩值
	ToPrice      float64 `json:"toPrice,omitempty"`      //单价最⼤值
	FromWlRate   int     `json:"fromWlRate,omitempty"`   //佣⾦⽐例最⼩值
	ToWlRate     int     `json:"toWlRate,omitempty"`     //佣⾦⽐例最⼤值
	//排序字段
	//(单价：wlPrice,
	//佣⾦⽐例： wlCommissionShare,
	//佣⾦： wlCommission，
	//30天 引单量：inOrderCount30Days，
	//30天⽀出佣⾦：inOrderComm30Days,
	//好评率：goodCommentsShare)
	SortName      string  `json:"sortName,omitempty"`
	Sort          string  `json:"sort,omitempty"`          //排序规则，升序还是降序（升序：asc，降序：desc）
	BrandCode     string  `json:"brandCode,omitempty"`     //品牌code N String
	ShopLevelFrom float64 `json:"shopLevelFrom,omitempty"` //最低店铺等级，⽀持枚举值：0，2.5，3.0，3.5，4.0，4.5，4.9
	JxFlags       []int   `json:"jxFlags,omitempty"`       //查询京喜商品，1:京喜商品，2：京喜⼯⼚ 3： List京喜优选
}

// GoodsQueryResp 商品查询返回数据
type GoodsQueryResp struct {
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		GoodsList []struct {
			AbParam        string `json:"ab_param"`       //推荐AB参数（默认推荐⻚及默认搜索⻚下发）
			AbVersion      string `json:"ab_version"`     // 推荐AB参数（默认推荐⻚及默认搜索⻚下发）
			Adowner        string `json:"adowner"`        //⼴告主
			BestCouponID   string `json:"bestCouponId"`   // 最优优惠券batchId，购买⼀件商品可⽤的最⼤⾯额优惠券
			BestCouponType int    `json:"bestCouponType"` // 最优优惠券类型1：通⽤券 2：计划券 3：提报券 4：公开券 5： 媒体券 6：⼆合⼀券
			BonusInfo      []struct {
				Id        int     `json:"id"`        // 活动ID
				Name      string  `json:"name"`      // 活动名称
				State     int     `json:"state"`     // 活动状态(1 已提报 2 预热中 3 进⾏中 4 已结束)
				BeginTime string  `json:"beginTime"` // 活动开始时间
				EndTime   string  `json:"endTime"`   // 活动结束时间
				Remark    string  `json:"remark"`    // 活动概述
				BonusType int     `json:"bonusType"` // 奖励⽅式 （ 0 定额 1 固定⽐例 2 提升⽐例）
				BonusRate float64 `json:"bonusRate"` // 奖励⽐例
				Level     int     `json:"level"`     // 活动等级（ 默认 0 共享 1 、2...）
				UdtList   []struct {
					Udt  string `json:"udt"`  // 渠道code
					Name string `json:"name"` // 渠道名字
				} `json:"udt_list"` //推⼴渠道（展示渠道udt-渠道code,name-名称 ，有序）
			} `json:"bonusInfo"` // 奖励活动列表（开发中）
			BrandCode string `json:"brandCode"` // 品牌code
			BrandName string `json:"brandName"` // 品牌名字
			Cid1      int    `json:"cid1"`      // ⼀级类⽬Id
			Cid1Name  string `json:"cid1Name"`  // ⼀级类⽬名称
			Cid2      int    `json:"cid2"`      // ⼆级类⽬Id
			Cid2Name  string `json:"cid2Name"`  // ⼆级类⽬名称
			Cid3      int    `json:"cid3"`      // 三级类⽬Id
			Cid3Name  string `json:"cid3Name"`  // 三级类⽬名称
			Color     string `json:"color"`     //颜⾊
			Comments  int    `json:"comments"`  // 商品评价数
			Coupons   []struct {
				AvailableTime  int64   `json:"availableTime"` //领取后可⽤时间
				BatchID        int     `json:"batchId"`       //批次id（唯⼀标识）
				BeginTime      int64   `json:"beginTime"`     //券领取开始时间
				CouponHotValue int64   `json:"couponHotValue"`
				CouponID       string  `json:"couponId"`
				CouponKind     int     `json:"couponKind"` //券种类
				CouponStyle    int     `json:"couponStyle"`
				CouponType     int     `json:"couponType"` //券的类型
				Discount       float64 `json:"discount"`   //优惠卷⾦额
				EndTime        int64   `json:"endTime"`
				Key            string  `json:"key"`
				Link           string  `json:"link"`         //券链接
				PlanID         int64   `json:"planId"`       //计划Id
				PlatformType   int     `json:"platformType"` //券限制使⽤平台
				Quota          float64 `json:"quota"`        //消费限额
				RemainCnt      int     `json:"remainCnt"`    //剩余券数量
				Source         string  `json:"source"`       //数据源
				Type           int     `json:"type"`         //券类型 1：通⽤，2：计划，3：提报，4：公开, 6:⼆合⼀转链
				UseEndTime     int64   `json:"useEndTime"`   //券使⽤结束时间
				UseStartTime   int64   `json:"useStartTime"` //券使⽤开始时间
				VenderID       int     `json:"venderId"`     //券的供应商id
			} `json:"coupons"` // 优惠券信息集合
			DeliveryType            int     `json:"deliveryType"`            // 京东配送 1：是，0：不是
			DiscountPrice           float64 `json:"discountPrice"`           // 券后价
			DiscountPriceCommission float64 `json:"discountPriceCommission"` // 券后价佣⾦
			EliteType               []int   `json:"eliteType"`               //精品库类型 1、算法、2、终结者、3、团⻓招商、4、内容商品
			EndTime                 int64   `json:"endTime"`                 //通⽤计划推⼴结束时间
			GoodComments            int     `json:"goodComments"`            //商品好评数
			GoodCommentsShare       int     `json:"goodCommentsShare"`       // 好评度
			HasCoupon               int     `json:"hasCoupon"`               // 是否有优惠券 0： 没有， 1： 有
			ImageURL                string  `json:"imageUrl"`                // 商品主图
			ImgList                 string  `json:"imgList"`                 // 轮播图集合
			InOrderComm30Days       float64 `json:"inOrderComm30Days"`       // 30天引⼊订单量
			InOrderCount30Days      int     `json:"inOrderCount30Days"`      // 30天引⼊佣⾦
			IsCare                  int     `json:"isCare"`                  //是否爱⼼咚咚商品 1：是，0：不是
			IsHot                   int     `json:"isHot"`                   // 是否为爆款 1：热⻔，0：不是热⻔
			IsLock                  int     `json:"isLock"`                  //是否锁定 0：没锁定， 1：锁定
			IsPinGou                int     `json:"isPinGou"`                // 是否可进⾏拼购推⼴（0或null为不拼购， 1 拼购）
			IsSeckill               int     `json:"isSeckill"`               //是否秒杀价 1：是 ， 0：不是
			JxFlags                 int     `json:"jxFlags"`                 //，1:京喜 商品，2：京喜⼯⼚ 3： List 京喜优选
			LowestPrice             float64 `json:"lowestPrice"`             // 最低价
			LowestPriceType         int     `json:"lowestPriceType"`         //最低价格类型1：⽆线价格 2：拼购价格 3：秒杀价格
			MajorSuppBrevityCode    string  `json:"majorSuppBrevityCode"`    // 供应商码
			MaterialID              string  `json:"materialId"`              // 商品落地⻚
			Orientation             []struct {
				Adowner       string `json:"adowner"`       //⼴告主
				BatchID       int    `json:"batchId"`       //优惠券批次id
				DisplayType   int    `json:"displayType"`   //是否展示
				DxType        int    `json:"dxType"`        //定向类型 0全部站⻓,1精确定向,2指定站⻓,3特殊媒体
				EndTime       string `json:"endTime"`       //计划结束时间
				IsHot         int    `json:"isHot"`         //是否爆品
				IsLock        int    `json:"isLock"`        //是否锁定佣⾦ 0：没锁定， 1：锁定
				MediaType     int    `json:"mediaType"`     //享橙和220V计划状态位
				PlanID        int64  `json:"planId"`        //计划id
				PlanStatus    int    `json:"planStatus"`    //计划状态
				PlanType      int    `json:"planType"`      //计划类型
				PromoteSource int    `json:"promoteSource"` //推⼴渠道 0联盟,1站内,2normal,3cps联盟活动,4拼购活动
				RuleType      int    `json:"ruleType"`      //规则类型
				StartTime     string `json:"startTime"`     //计划开始时间
				WlRatio       int    `json:"wlRatio"`       //佣⾦⽐例
			} `json:"orientation"` //商品定向信息集合
			OrientationFlag int     `json:"orientationFlag"` // 是否有定向计划 1：有，0：没有
			Owner           string  `json:"owner"`           //g为⾃营 、p为pop
			Pid             int64   `json:"pid"`             // 主商品id，spu维度
			PingouActiveID  int64   `json:"pingouActiveId"`  // 拼购活动Id
			PingouEnd       int64   `json:"pingouEnd"`       // 拼购结束时间
			PingouPrice     float64 `json:"pingouPrice"`     //拼购价格
			PingouStart     int64   `json:"pingouStart"`     //拼购开始时间
			PingouTagList   []struct {
				tagName string //标签名字
			} `json:"pingouTagList"` //拼购信息集合
			PingouTmCount      int     `json:"pingouTmCount"`      // 拼购成团所需⼈数
			PlanID             int64   `json:"planId"`             //通⽤计划id
			QualityScore       float64 `json:"qualityScore"`       // 商品打分
			RequestID          string  `json:"requestId"`          //搜索打标
			RuleType           int64   `json:"ruleType"`           //规则类型
			SeckillDisplayTime int64   `json:"seckillDisplayTime"` // 秒杀展示时间
			SeckillEndTime     int64   `json:"seckillEndTime"`     // 秒杀结束时间
			SeckillOriPrice    float64 `json:"seckillOriPrice"`    // 秒杀价原价
			SeckillPrice       float64 `json:"seckillPrice"`       //秒杀价格
			SeckillStartTime   int64   `json:"seckillStartTime"`   // 秒杀开始时间
			ShopID             int     `json:"shopId"`             //店铺id
			Size               string  `json:"size"`               // 型号
			SkuID              int64   `json:"skuId"`              //商品skuId
			SkuName            string  `json:"skuName"`            // 商品名称
			StartTime          int64   `json:"startTime"`          // 通⽤计划推⼴开始时间
			VenderName         string  `json:"venderName"`         // 供应商
			Vid                int     `json:"vid"`                // 商家id
			WareID             int64   `json:"wareId"`             // 款id
			WlCommission       float64 `json:"wlCommission"`       // ⽆线佣⾦
			WlCommissionShare  float64 `json:"wlCommissionShare"`  //⽆线佣⾦⽐例
			WlPrice            float64 `json:"wlPrice"`            // 单价
		} `json:"goodsList"`
		PageNo     int `json:"pageNo"`
		PageSize   int `json:"pageSize"`
		TotalItems int `json:"totalItems"`
	} `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
}

// PromotionGetParams 获取推荐链接参数
type PromotionGetParams struct {
	Code          string `json:"code"`                    //渠道编码，京佣系统分配
	MaterialID    string `json:"materialId"`              //落地⻚地址，也可以是活动⻚地址
	CouponURL     string `json:"couponUrl,omitempty"`     //优惠券地址
	SubUnionID    string `json:"subUnionId,omitempty"`    //⼦联盟Id，每个下级必须唯⼀，不要有中⽂,不要包含逗号","，不要以partner开头，⻓度低于40
	EuID          string `json:"euId,omitempty"`          //扩展字段，⻓度不超过30
	OriginalURL   string `json:"originalUrl,omitempty"`   //使⽤原始推⼴链接⽣成新的推⼴链接(替 换落地⻚，取此链接中的站⻓信息)
	SourceEmt     string `json:"sourceEmt,omitempty"`     //推⼴媒介
	RequestID     string `json:"requestId,omitempty"`     //搜索来源
	GiftCouponKey string `json:"giftCouponKey,omitempty"` //礼⾦加密key 当传⼊礼⾦批次ID时，materialId必须 传对应的商品落地⻚
}

// PromotionGetResp 获取推荐链接返回数据
type PromotionGetResp struct {
	Code    int    `json:"code"`
	HasNext bool   `json:"hasNext"`
	Message string `json:"message"`
	Result  struct {
		ClickURL string `json:"clickUrl"` //长链接
		ShortURL string `json:"shortUrl"` //短链接
	} `json:"result"`
	TotalNum int         `json:"totalNum"`
	UUID     interface{} `json:"uuid"`
}

// OrderQueryParams 订单查询参数
type OrderQueryParams struct {
	Code      string `json:"code"`      //渠道编码，京佣系统分配，联系商务申请
	StartTime string `json:"startTime"` //开始时间（含），格式yyyy-MM-dd HH:mm:ss（entTime和startTime时间间隔最多30 分钟）
	EndTime   string `json:"endTime"`   //结束时间（不含），格式yyyy-MM-dd HH:mm:ss（entTime和startTime时间间隔最多30 分钟）
	PageSize  int    `json:"pageSize"`  //每⻚记录数，⽬前最多500条
	PageNo    int    `json:"pageNo"`    //查询的当前页
	Type      int    `json:"type"`      //1：下单时间，2：完成时间，3：更新时间，默认 下单时间
}

// OrderQueryResp 订单查询返回数据
type OrderQueryResp struct {
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		HasNext  bool `json:"hasNext"` //是否有下一页
		PageNo   int  `json:"pageNo"`
		PageSize int  `json:"pageSize"`
		Result   []struct {
			ActualCommission   float64 `json:"actualCommission"`   //sku⾏总的实际预估佣⾦
			ActualCosPrice     float64 `json:"actualCosPrice"`     //实际计佣⾦额
			ActualFee          float64 `json:"actualFee"`          //实际预估付给推⼴者的费⽤
			CommissionRate     float64 `json:"commissionRate"`     //佣⾦⽐例
			CpActID            int     `json:"cpActId"`            //招商团活动id
			EstimateCommission float64 `json:"estimateCommission"` //sku⾏总的预估佣⾦
			EstimateCosPrice   float64 `json:"estimateCosPrice"`   //预估计佣⾦额
			EstimateFee        float64 `json:"estimateFee"`        //预估付给推⼴者的费⽤
			FinalRate          float64 `json:"finalRate"`          //最终⽐例 (⼀级分佣⽐例*⼆级分佣⽐例)
			FinishTime         int64   `json:"finishTime"`         //完成时间
			FirstLevel         int     `json:"firstLevel"`         //⼀级类⽬
			FrozenSkuNum       int     `json:"frozenSkuNum"`       //商品售后中数量
			OrderID            int64   `json:"orderId"`            //订单号
			OrderTime          int64   `json:"orderTime"`          //下单时间
			ParentID           int64   `json:"parentId"`           //⽗单号，在被拆单的情况下会有
			PayMonth           int64   `json:"payMonth"`           //结算时间
			PayPrice           float64 `json:"payPrice"`           //实际⽀付⾦额
			Pid                string  `json:"pid"`                //pid格式:⼦站⻓ID⼦站⻓⽹站ID⼦站⻓推⼴位ID
			Price              float64 `json:"price"`              //商品单价
			RowID              int64   `json:"rowId"`              //分配⾏标识
			SecondLevel        int     `json:"secondLevel"`        //⼆级类⽬
			SiteID             int64   `json:"siteId"`             //⽹站ID
			SkuID              int64   `json:"skuId"`              //商品SKU
			SkuName            string  `json:"skuName"`            //商品名称
			SkuNum             int     `json:"skuNum"`             //商品下单数量
			SkuReturnNum       int     `json:"skuReturnNum"`       //商品退货数量
			SpID               int64   `json:"spId"`               //推⼴位ID
			SubSideRate        float64 `json:"subSideRate"`        //⼀级分成⽐例
			SubUnionID         string  `json:"subUnionId"`         //⼦联盟ID
			SubsidyRate        float64 `json:"subsidyRate"`        //⼀级补贴⽐例
			ThirdLevel         int     `json:"thirdLevel"`         //三级类⽬
			TraceType          int     `json:"traceType"`          //同跨店：0影响 1直接2同店 3跨店
			UnionAlias         string  `json:"unionAlias"`         //⺟账号简称
			UnionRole          int     `json:"unionRole"`          //站⻓⻆⾊：1 推客 2 团⻓
			UnionTag           string  `json:"unionTag"`           //联盟标签数据
			UnionTrafficGroup  int     `json:"unionTrafficGroup"`  //渠道组 (1, "1号店", "1号店") (2, "开普勒", "京东") (3, "超 级返", "京东") (4, "京粉", "京东") (5, "站⻓media", "京 东")
			//联盟订单状态 -1未知 2-14⽆效 15待付 16已付 17已完 成 18已结算-废弃，具体含义⻅下表
			//1 未知
			//2 ⽆效-拆单
			//3 ⽆效-取消
			//4 ⽆效-京东帮帮主订单
			//5 ⽆效-账号异常
			//6 ⽆效-赠品类⽬不返佣
			//7 ⽆效-校园订单
			//8 ⽆效-企业订单
			//9 ⽆效-团购订单
			//10 ⽆效-开增值税专⽤发票订单
			//11 ⽆效-乡村推⼴员下单
			//12 ⽆效-⾃⼰推⼴⾃⼰下单
			//13 ⽆效-违规订单
			//14 ⽆效-来源与备案⽹址不符
			//15 待付款
			//16 已付款
			//17 已完成
			//18 已结算，该状态值将被废弃
			//19 ⽆效-佣⾦⽐例为0（部分站⻓会返回这个值，由业务运营确定站⻓名单）
			ValidCode int `json:"validCode"`
		} `json:"result"`
		TotalItems int `json:"totalItems"` //符合条件的总数量
	} `json:"data"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
}
