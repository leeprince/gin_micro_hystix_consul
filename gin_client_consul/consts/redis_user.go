package consts

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/19 下午10:41
 * @Desc:
 */

const USER_KEY string = "user:%d"
// 单位：秒(time.Duration:纳秒), 1e9 = 1秒
const USER_EXPIE time.Duration = 1e9 * 120