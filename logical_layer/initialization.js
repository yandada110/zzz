/* initialization.js */

// 定义常量（全局可用）
const ATTACK_PERCENTAGE = "AttackPercentage";
const CRITICAL = "Critical";
const EXPLOSIVE_INJURY = "ExplosiveInjury";
const INCREASED_DAMAGE = "IncreasedDamage";

// 基础类定义
class Basic {
    constructor(basicAttack, basicCritical, basicExplosiveInjury, basicIncreasedDamage, basicReductionResistance, basicVulnerable, basicSpecialDamage) {
        this.basicAttack = basicAttack;
        this.basicCritical = basicCritical;
        this.basicExplosiveInjury = basicExplosiveInjury;
        this.basicIncreasedDamage = basicIncreasedDamage;
        this.basicReductionResistance = basicReductionResistance;
        this.basicVulnerable = basicVulnerable;
        this.basicSpecialDamage = basicSpecialDamage;
    }
}

class Gain {
    constructor(attackValue, attackPowerPercentage, attackInternalPercentage, critical, explosiveInjury, increasedDamage, reductionResistance, vulnerable, specialDamage) {
        this.attackValue = attackValue;
        this.attackPowerPercentage = attackPowerPercentage;
        this.attackInternalPercentage = attackInternalPercentage;
        this.critical = critical;
        this.explosiveInjury = explosiveInjury;
        this.increasedDamage = increasedDamage;
        this.reductionResistance = reductionResistance;
        this.vulnerable = vulnerable;
        this.specialDamage = specialDamage;
    }
}

class Defense {
    constructor(penetration, defenseBreak, penetrationValue) {
        this.penetration = penetration;
        this.defenseBreak = defenseBreak;
        this.penetrationValue = penetrationValue;
    }
}

class Condition {
    constructor(mainArticle, critical) {
        this.mainArticle = mainArticle;
        this.critical = critical;
        this.criticalStatus = 0;
        this.criticalCount = 0;
    }
}

class Output {
    constructor() {
        this.basicDamageArea = 0;
        this.increasedDamageArea = 0;
        this.explosiveInjuryArea = 0;
        this.defenseArea = 0;
        this.reductionResistanceArea = 0;
        this.vulnerableArea = 0;
        this.specialDamageArea = 0;
    }
}

class CurrentPanel {
    constructor(attack = 0, critical = 0, explosiveInjury = 0, increasedDamage = 0, reductionResistance = 0, vulnerable = 0, specialDamage = 0) {
        this.attack = attack;
        this.critical = critical;
        this.explosiveInjury = explosiveInjury;
        this.increasedDamage = increasedDamage;
        this.reductionResistance = reductionResistance;
        this.vulnerable = vulnerable;
        this.specialDamage = specialDamage;
    }
}

class Magnification {
    constructor(magnificationValue, triggerTimes, name, increasedDamage = 0, reductionResistance = 0, defenseBreak = 0, penetration = 0, specialDamage = 0) {
        this.magnificationValue = magnificationValue;
        this.triggerTimes = triggerTimes;
        this.name = name;
        this.increasedDamage = increasedDamage;
        this.reductionResistance = reductionResistance;
        this.defenseBreak = defenseBreak;
        this.penetration = penetration;
        this.specialDamage = specialDamage;
    }
}

// Initialization 类：负责初始化当前面板及基础属性更新
class Initialization {
    constructor(basic, gain, defense, condition, magnifications) {
        this.basic = basic;
        this.gain = gain;
        this.defense = defense;
        this.condition = condition;
        this.magnifications = magnifications;
        this.output = new Output();
        // 初始化当前面板（部分数值为 Basic 与 Gain 累加）
        this.currentPanel = new CurrentPanel();
        this.currentPanel.reductionResistance = this.basic.basicReductionResistance + this.gain.reductionResistance;
        this.currentPanel.vulnerable = this.basic.basicVulnerable + this.gain.vulnerable;
        this.currentPanel.specialDamage = this.basic.basicSpecialDamage + this.gain.specialDamage;
        // 初始调用：未分配词条
        this.handleBasicAttack("", 0);
        this.handleBasicCritical("", 0);
        this.handleBasicExplosiveInjury("", 0);
        this.handleBasicIncreasedDamage("", 0);
    }

    // 根据指定词条和数量更新面板
    characterPanel(key, count) {
        this.currentPanel = new CurrentPanel();
        this.currentPanel.reductionResistance = this.basic.basicReductionResistance + this.gain.reductionResistance;
        this.currentPanel.vulnerable = this.basic.basicVulnerable + this.gain.vulnerable;
        this.currentPanel.specialDamage = this.basic.basicSpecialDamage + this.gain.specialDamage;
        this.handleBasicAttack(key, count);
        this.handleBasicExplosiveInjury(key, count);
        this.handleBasicCritical(key, count);
        this.handleBasicIncreasedDamage(key, count);
    }

    handleBasicAttack(key, count) {
        let attackPowerPercentage = this.gain.attackPowerPercentage;
        if (key === ATTACK_PERCENTAGE) {
            attackPowerPercentage += 3 * count;
        }
        this.currentPanel.attack = (this.basic.basicAttack * (1 + attackPowerPercentage / 100) + this.gain.attackValue) *
            (1 + this.gain.attackInternalPercentage / 100);
    }

    handleBasicCritical(key, count) {
        if (key === CRITICAL) {
            let crit = this.basic.basicCritical + this.gain.critical + 2.4 * count;
            if (crit > this.condition.critical) {
                this.condition.criticalStatus++;
                this.currentPanel.critical = this.basic.basicCritical + this.gain.critical + 2.4 * (count - this.condition.criticalStatus);
                this.currentPanel.explosiveInjury = this.currentPanel.explosiveInjury + this.condition.criticalStatus * 4.8;
            } else {
                this.condition.criticalCount = count;
                this.currentPanel.critical = crit;
            }
        } else {
            this.currentPanel.critical = this.basic.basicCritical + this.gain.critical;
        }
    }

    handleBasicExplosiveInjury(key, count) {
        let explosiveInjury = this.gain.explosiveInjury;
        if (key === EXPLOSIVE_INJURY) {
            explosiveInjury += 4.8 * count;
        }
        this.currentPanel.explosiveInjury = this.basic.basicExplosiveInjury + explosiveInjury;
    }

    handleBasicIncreasedDamage(key, count) {
        let increasedDamage = this.gain.increasedDamage;
        if (key === INCREASED_DAMAGE) {
            increasedDamage += 3 * count;
        }
        this.currentPanel.increasedDamage = this.basic.basicIncreasedDamage + increasedDamage;
    }
}

// UI 工具函数：添加/删除 Magnification 项
function addMagnification(defaults) {
    defaults = defaults || {
        magnificationValue: 943.5,
        triggerTimes: 2,
        name: "普攻",
        increasedDamage: 0,
        reductionResistance: 0,
        defenseBreak: 0,
        penetration: 0,
        specialDamage: 0
    };
    const div = document.createElement("div");
    div.className = "magnification-item";
    div.innerHTML = `
        <fieldset style="border:1px solid #ccc; padding:10px; margin-bottom:10px;">
          <legend>Magnification</legend>
          <label>名称: <input type="text" name="magName" value="${defaults.name}"></label><br>
          <label>倍率值: <input type="number" step="any" name="magValue" value="${defaults.magnificationValue}"></label><br>
          <label>触发次数: <input type="number" step="any" name="magTrigger" value="${defaults.triggerTimes}"></label><br>
          <label>增伤: <input type="number" step="any" name="magIncreasedDamage" value="${defaults.increasedDamage}"></label><br>
          <label>减抗: <input type="number" step="any" name="magReductionResistance" value="${defaults.reductionResistance}"></label><br>
          <label>破防: <input type="number" step="any" name="magDefenseBreak" value="${defaults.defenseBreak}"></label><br>
          <label>穿透: <input type="number" step="any" name="magPenetration" value="${defaults.penetration}"></label><br>
          <label>特殊增伤: <input type="number" step="any" name="magSpecialDamage" value="${defaults.specialDamage}"></label><br>
          <button type="button" onclick="removeMagnification(this)">删除</button>
        </fieldset>
    `;
    document.getElementById("magnifications").appendChild(div);
}

function removeMagnification(button) {
    const item = button.parentElement.parentElement;
    item.parentElement.removeChild(item);
}

// 页面加载时初始化默认倍率项
window.onload = function() {
    addMagnification({magnificationValue: 943.5, triggerTimes: 2, name: "普攻", increasedDamage: 0, reductionResistance: 0, defenseBreak: 0, penetration: 0, specialDamage: 0});
    addMagnification({magnificationValue: 1658.7, triggerTimes: 4.5, name: "连携技", increasedDamage: 30, reductionResistance: 25, defenseBreak: 0, penetration: 0, specialDamage: 25});
    addMagnification({magnificationValue: 1202.5, triggerTimes: 2, name: "强化特殊技", increasedDamage: 0, reductionResistance: 0, defenseBreak: 0, penetration: 0, specialDamage: 0});
    addMagnification({magnificationValue: 3977.3, triggerTimes: 1, name: "终结技", increasedDamage: 30, reductionResistance: 25, defenseBreak: 0, penetration: 0, specialDamage: 25});
};
