/* attributeCalculation.js */

// 工具函数：计算百分比差值
function DecimalToPercentage(newNumber, oldNumber) {
    if (oldNumber === 0) return 0;
    let decimalPart = (newNumber - oldNumber) / oldNumber;
    let percentage = decimalPart * 100;
    return Math.floor(percentage * 1000 + 0.5) / 1000;
}

// 分配逻辑：循环分配各词条（挂在 Initialization 原型上）
Initialization.prototype.getMaxFloat = function(mainArticleDamageMap) {
    let log = "";
    let entriesMap = {};
    entriesMap[ATTACK_PERCENTAGE] = 0;
    entriesMap[CRITICAL] = 0;
    entriesMap[EXPLOSIVE_INJURY] = 0;
    entriesMap[INCREASED_DAMAGE] = 0;
    entriesMap[CRITICAL + "--"] = 0;
    let criticalStatus = false;
    // 循环次数为 condition.mainArticle
    for (let a = 0; a < this.condition.mainArticle; a++) {
        let optimal = CRITICAL;
        let optimalKey = CRITICAL + "--";
        // 与 攻击力百分比 比较
        if ( mainArticleDamageMap[optimal][ entriesMap[optimalKey] ].PercentageDifference <
            mainArticleDamageMap[ATTACK_PERCENTAGE][ entriesMap[ATTACK_PERCENTAGE] ].PercentageDifference ) {
            optimal = ATTACK_PERCENTAGE;
            optimalKey = ATTACK_PERCENTAGE;
        }
        // 与 爆伤 比较
        if ( mainArticleDamageMap[optimal][ entriesMap[optimalKey] ].PercentageDifference <
            mainArticleDamageMap[EXPLOSIVE_INJURY][ entriesMap[EXPLOSIVE_INJURY] ].PercentageDifference ) {
            optimal = EXPLOSIVE_INJURY;
            optimalKey = EXPLOSIVE_INJURY;
        }
        // 与 增伤 比较
        if ( mainArticleDamageMap[optimal][ entriesMap[optimalKey] ].PercentageDifference <
            mainArticleDamageMap[INCREASED_DAMAGE][ entriesMap[INCREASED_DAMAGE] ].PercentageDifference ) {
            optimal = INCREASED_DAMAGE;
            optimalKey = INCREASED_DAMAGE;
        }
        let explosiveInjuryStatus = false;
        if (optimal === CRITICAL) {
            entriesMap[CRITICAL + "--"]++;
            if (criticalStatus) {
                explosiveInjuryStatus = true;
                optimal = EXPLOSIVE_INJURY;
                optimalKey = EXPLOSIVE_INJURY;
            } else {
                if (this.condition.criticalCount === entriesMap[CRITICAL]) {
                    explosiveInjuryStatus = true;
                    criticalStatus = true;
                    optimal = EXPLOSIVE_INJURY;
                    optimalKey = EXPLOSIVE_INJURY;
                }
            }
        }
        // 记录本次增加的词条数
        entriesMap[optimal] = (entriesMap[optimal] || 0) + 1;
        if (explosiveInjuryStatus) {
            log += `<br>-----------------------------------------------------------------------------------<br>词条分配：攻击力：${entriesMap[ATTACK_PERCENTAGE]}个，暴击：${entriesMap[CRITICAL]}个，爆伤：${entriesMap[EXPLOSIVE_INJURY]}个，增伤：${entriesMap[INCREASED_DAMAGE]}个 ` +
                `差距：${mainArticleDamageMap[CRITICAL][ entriesMap[CRITICAL + "--"] ].PercentageDifference.toFixed(6)} 总伤：${mainArticleDamageMap[CRITICAL][ entriesMap[CRITICAL + "--"] ].Damage.toFixed(6)}<br>`;
        } else {
            log += `<br>----------------------------------------------------------------------------------<br>词条分配：攻击力：${entriesMap[ATTACK_PERCENTAGE]}个，暴击：${entriesMap[CRITICAL]}个，爆伤：${entriesMap[EXPLOSIVE_INJURY]}个，增伤：${entriesMap[INCREASED_DAMAGE]}个 ` +
                `差距：${mainArticleDamageMap[optimal][ entriesMap[optimal] ].PercentageDifference.toFixed(6)} 总伤：${mainArticleDamageMap[optimal][ entriesMap[optimal] ].Damage.toFixed(6)}<br>`;
        }
    }
    // 更新最终面板
    this.currentPanel.reductionResistance = this.basic.basicReductionResistance + this.gain.reductionResistance;
    this.currentPanel.vulnerable = this.basic.basicVulnerable + this.gain.vulnerable;
    this.currentPanel.specialDamage = this.basic.basicSpecialDamage + this.gain.specialDamage;
    this.handleBasicAttack(ATTACK_PERCENTAGE, entriesMap[ATTACK_PERCENTAGE] || 0);
    this.handleBasicCritical(CRITICAL, entriesMap[CRITICAL] || 0);
    this.handleBasicExplosiveInjury(EXPLOSIVE_INJURY, entriesMap[EXPLOSIVE_INJURY] || 0);
    this.handleBasicIncreasedDamage(INCREASED_DAMAGE, entriesMap[INCREASED_DAMAGE] || 0);
    for (let mag of this.magnifications) {
        this.increasedDamageArea(mag);
        this.reductionResistanceArea(mag);
    }
    this.vulnerableArea();
    this.specialDamageArea();
    log += `<br>最终面板：<br>`;
    log += `  攻击力: ${this.currentPanel.attack.toFixed(2)}, 暴击: ${this.currentPanel.critical.toFixed(2)}%, `;
    log += `爆伤: ${this.currentPanel.explosiveInjury.toFixed(2)}%, 增伤: ${((this.output.increasedDamageArea - 1) * 100).toFixed(2)}%<br>`;
    log += `  抗性区: ${(((this.output.reductionResistanceArea - 1) * 100)).toFixed(2)}%, `;
    log += `易伤区: ${(((this.output.vulnerableArea - 1) * 100)).toFixed(2)}%, `;
    log += `特殊乘区: ${(((this.output.specialDamageArea - 1) * 100)).toFixed(2)}%<br>`;
    return log;
};

// 主函数：读取表单数据，循环各属性计算伤害并分配词条
function runCalculation() {
    let logMessages = "";
    const form = document.getElementById("calculatorForm");
    // 读取 Basic 参数
    const basic = new Basic(
        parseFloat(form.basicAttack.value) || 0,
        parseFloat(form.basicCritical.value) || 0,
        parseFloat(form.basicExplosiveInjury.value) || 0,
        parseFloat(form.basicIncreasedDamage.value) || 0,
        parseFloat(form.basicReductionResistance.value) || 0,
        parseFloat(form.basicVulnerable.value) || 0,
        parseFloat(form.basicSpecialDamage.value) || 0
    );
    // 读取 Gain 参数
    const gain = new Gain(
        parseFloat(form.attackValue.value) || 0,
        parseFloat(form.attackPowerPercentage.value) || 0,
        parseFloat(form.attackInternalPercentage.value) || 0,
        parseFloat(form.gainCritical.value) || 0,
        parseFloat(form.gainExplosiveInjury.value) || 0,
        parseFloat(form.gainIncreasedDamage.value) || 0,
        parseFloat(form.gainReductionResistance.value) || 0,
        parseFloat(form.gainVulnerable.value) || 0,
        parseFloat(form.gainSpecialDamage.value) || 0
    );
    // 读取 Defense 参数
    const defense = new Defense(
        parseFloat(form.penetration.value) || 0,
        parseFloat(form.defenseBreak.value) || 0,
        parseFloat(form.penetrationValue.value) || 0
    );
    // 读取 Condition 参数
    const condition = new Condition(
        parseFloat(form.mainArticle.value) || 0,
        parseFloat(form.conditionCritical.value) || 0
    );
    // 读取 Magnifications 数组
    const magDivs = document.querySelectorAll(".magnification-item");
    let magnifications = [];
    magDivs.forEach(div => {
        const name = div.querySelector("input[name='magName']").value;
        const magValue = parseFloat(div.querySelector("input[name='magValue']").value) || 0;
        const triggerTimes = parseFloat(div.querySelector("input[name='magTrigger']").value) || 0;
        const increasedDamage = parseFloat(div.querySelector("input[name='magIncreasedDamage']").value) || 0;
        const reductionResistance = parseFloat(div.querySelector("input[name='magReductionResistance']").value) || 0;
        const defenseBreak = parseFloat(div.querySelector("input[name='magDefenseBreak']").value) || 0;
        const penetration = parseFloat(div.querySelector("input[name='magPenetration']").value) || 0;
        const specialDamage = parseFloat(div.querySelector("input[name='magSpecialDamage']").value) || 0;
        magnifications.push(new Magnification(magValue, triggerTimes, name, increasedDamage, reductionResistance, defenseBreak, penetration, specialDamage));
    });

    // 构造 Initialization 对象
    const init = new Initialization(basic, gain, defense, condition, magnifications);
    // 构造记录每种词条类型计算结果的 Map
    let mainArticleDamageMap = {};
    mainArticleDamageMap[ATTACK_PERCENTAGE] = [];
    mainArticleDamageMap[CRITICAL] = [];
    mainArticleDamageMap[EXPLOSIVE_INJURY] = [];
    mainArticleDamageMap[INCREASED_DAMAGE] = [];
    // 初始化各类型当前词条分配数量
    let MainArticleMap = {};
    MainArticleMap[ATTACK_PERCENTAGE] = 0;
    MainArticleMap[CRITICAL] = 0;
    MainArticleMap[EXPLOSIVE_INJURY] = 0;
    MainArticleMap[INCREASED_DAMAGE] = 0;

    // 对每个词条类型分别进行计算
    const attributes = [ATTACK_PERCENTAGE, CRITICAL, EXPLOSIVE_INJURY, INCREASED_DAMAGE];
    attributes.forEach(key => {
        let attributeName = "";
        switch (key) {
            case ATTACK_PERCENTAGE:
                attributeName = "攻击力";
                break;
            case CRITICAL:
                attributeName = "暴击率";
                break;
            case EXPLOSIVE_INJURY:
                attributeName = "爆伤";
                break;
            case INCREASED_DAMAGE:
                attributeName = "增伤";
                break;
        }
        logMessages += `<br>正在计算属性：${attributeName}<br>`;
        logMessages += `------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------<br>`;
        logMessages += `------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------<br>`;
        // 初始面板
        init.characterPanel("", 0);
        let OldDamage = init.calculatingTotalDamage();
        MainArticleMap[key] = 1;
        while (MainArticleMap[key] <= condition.mainArticle) {
            init.characterPanel(key, MainArticleMap[key]);
            logMessages += `当前面板属性：第${MainArticleMap[key]}词条分配， 攻击力: ${init.currentPanel.attack.toFixed(2)}, 暴击: ${init.currentPanel.critical.toFixed(2)}%, `;
            logMessages += `爆伤: ${init.currentPanel.explosiveInjury.toFixed(2)}%, 增伤: ${init.currentPanel.increasedDamage.toFixed(2)}%<br>`;
            let NewDamage = init.calculatingTotalDamage();
            let percentageDifference = DecimalToPercentage(NewDamage, OldDamage);
            mainArticleDamageMap[key].push({
                CurrentPanel: Object.assign({}, init.currentPanel),
                Output: Object.assign({}, init.output),
                Damage: NewDamage,
                PercentageDifference: percentageDifference
            });
            logMessages += `输出总伤害: ${(NewDamage / 10000).toFixed(6)}万<br>`;
            logMessages += `伤害差距: ${percentageDifference.toFixed(2)}%<br>`;
            logMessages += `------------------------------------------------------<br>`;
            OldDamage = NewDamage;
            MainArticleMap[key]++;
        }
    });
    // 调用分配逻辑并获得最终日志
    logMessages += init.getMaxFloat(mainArticleDamageMap);
    document.getElementById("logOutput").innerHTML = logMessages;
}
