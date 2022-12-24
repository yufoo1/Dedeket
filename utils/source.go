package utils

import (
	"Dedeket/model/deal"
	"time"
)

func InitSource() {
	var bookNameArr = []string{"数学分析中的典型问题与方法", "普通物理学（下册）", "高等代数习题解", "概率统计与随机过程习题解集", "运筹学", "统计学习方法", "计算机组成与实现", "经济管理基础", "法理学", "生物细胞学（第4版）", "密码学中的可证明安全性", "数据库系统概论（第五版）", "形式语言与自动机理论（第3版）", "思想道德修养与法律基础", "组合数学"}
	var writerArr = []string{"裴礼文", "程守洙，江之永　主编", "杨子胥", "邢家省", "《运筹学》教材编写组", "李航", "高小鹏", "单大明", "时显群", "翟中和、王喜忠、丁明孝", "杨波", "王珊/萨师煊", "蒋宗礼、姜守旭", "《思想道德修养与法律基础》编写组", "杨雅琴、李秋月、马腾宇"}
	var classArr = []string{"数学分析", "普通物理学", "高等代数", "概率论", "计算机工程中的最优化问题", "机器学习", "计算机组成原理", "经济管理学", "法理学", "生物细胞学", "可证明安全", "数据库原理", "形式语言与自动机", "思想道德修养与法律基础", "组合数学"}
	var descriptionArr = []string{"遵循现行教材的顺序，本书全面、系统地总结和归纳了数学分析问题的基本类型，每种类型的基本方法，对每种方法先概括要点，再选取典型而有相当难度的例题，逐层剖析，分类讲解。然后分别配备相应的一套练习，旨在拓宽基础，启发思路了，培养学生分析问题和解决问题的能力，作教材的补充和延深。此外，对现行教材中比较薄弱的部分，如半连续、凸函数、不等式、等度连续等内容，作了适当扩充。选题具有很强的典型性、灵活性、启发性、趣味性和综合性，对培养学生的能力极为有益，可供数学类各专业师生及有关读者参考，也可供数学一的考生选择阅读。",
		"写于20 世纪50 年代的《普通物理学》(程守洙、江之永主编)是我国工科物理最早的教材之一，半个多世纪来已有多次修订和改编，在重要历史时期为我国高等学校大学物理课程的教学作出积极贡献.本文从历史与文化、基础与前沿、理论与实际等诸方面评述了该教材的改编特色，得出结论：此书是一本历史悠久、不断进取且能全面体现传输知识、培养能力和提高科学素质的大学物理教材。",
		"对高等代数解题方法进行总结，对高等代数题目进行详细的解析。最近刷完了，速出一本。",
		"《概率统计与随机过程习题解集》是《概率统计与随机过程》的习题解集，适用于理工科大学学生的学习。《概率统计与随机过程习题解集》对概率统计与随机过程中的常规性练习题目给出了解答，题型多样，覆盖面较全。通过练习和对照使用，有助于学生巩固已学的知识和理论，掌握解决基本问题的方法和手段，提高解决问题的能力，以期能熟练灵活地解决更多的问题，取到较好的效果。",
		"本书在修订版基础上，吸收了广大读者的意见，作了局部调整和修改。除原有线性规划、整数规划、非线性规划、动态规划、图与网络分析、排队论、存储论、对策论、决策论、目标规划和多目标决策以外，增加了启发式方法一章",
		"本书全面系统地介绍了统计学习的主要方法，共分两篇。第一篇系统介绍监督学习的各种重要方法，包括决策树、感知机、支持向量机、最大熵模型与逻辑斯谛回归、推进法、多类分类法、EM算法、隐马尔科夫模型和条件随机场等；第二篇介绍无监督学习，包括聚类、奇异值、主成分分析、潜在语义分析等。两篇中，除概论和总结外，每章介绍一或二种方法。",
		"本书为“基于系统能力培养的计算机专业课程建设研究”项目规划教材，根据北京航空航天大学对计算机专业系统能力培养十二年改革成果编写而成。本课程的教学目标为：开发一个具有数十条指令规模且通过严格测试的功能型CPU。",
		"《高等学校通识课程系列教材:经济管理基础》主要内容包括：导论、消费者行为理论、生产与成本、国民经济核算、财政与金融、WTO基本知识、企业与管理、企业组织与环境、企业经营管理、人力资源管理与组织行为、市场营销管理、生产运作管理、质量管理、物流管理、技术管理、会计基础、财务管理、统计基础。",
		"法理学又称为法哲学，是法学的一门基础学科，主要研究法律现象的共同规律和共同性问题。法理学的研究对象是一般法，但它的内容不是法的全部，而仅仅是包含在法律现象中的普遍问题和根本问题。",
		"《细胞生物学（第4版）》共17章，内容包括绪论、细胞的统一性与多样性、细胞生物学研究方法、细胞质膜、物质的跨膜运输、线粒体、叶绿体、细胞质基质与内膜系统、蛋白质分选与膜泡运输、细胞信号转导、细胞骨架、细胞核与染色质、核糖体、细胞周期与细胞分裂、细胞增殖调控与癌细胞、细胞分化与胚胎发育、细胞死亡、细胞衰老、细胞的社会联系等。",
		"本书全面介绍可证明安全性的发展历史及研究成果。全书共5章，第1章介绍可证明安全性涉及的数学知识和基本工具，第2章介绍语义安全的公钥密码体制的定义，第3章介绍几类常用的语义安全的公钥机密体制，第4章介绍基于身份的密码体制，第5章介绍基于属性的密码体制。",
		"全书分为4篇16章。第一篇基础篇，包括绪论、关系数据库、关系数据库标准语言SQL、数据库安全性和数据库完整性，共5章；第二篇设计与应用开发篇，包括关系数据理论、数据库设计和数据库编程，共3章；第三篇系统篇，包括关系查询处理和查询优化、数据库恢复技术、并发控制和数据库管理系统，共4章；第四篇新技术篇，包括数据库技术发展概述、大数据管理、内存数据库系统和数据仓库与联机分析处理技术，共4章",
		"本书是作者结合其近30年来在大学讲授该门课程的经验和体会，选择和组织有关内容撰写而成。基于计算机问题求解的需要讨论正则语言、上下文无关语言的文法、识别模型及其性质、图灵机的基本知识。其内容特点是抽象和形式化，既有严格的理论证明，又具有很强的构造性。叙述中特别注意引导读者分析与解决问题，以培养学生的形式化描述和抽象思维能力",
		"为了推动党的十八大精神进教材、进课堂、进头脑，体现上次修订以来中国特色社会主义理论和实践的创新成果，体现思想政治教育学科的新进展，中宣部、教育部组织课题组在广泛调研的基础上，再次对教材进行了修订。马克思主义理论研究和建设工程咨询委员会对教材修订稿进行了审议指导。",
		"本书作者多年教学和研究成果的基础上结合组合数学的基本理论，系统地介绍了组合计数、组合设计以及相关数学理论。全书分为11章，介绍了简单排列组合与多重集的简单排列组合、鸽巢原理和Ramsey(拉姆齐)定理、容斥原理、生成函数、递推方程、特殊计数、Burnside(伯恩赛德)定理和Pólya(波利亚)定理、图论、区组设计、编码理论等内容。"}
	var sellerArr = []string{"教务处", "教材处", "cyw1", "yufoo1", "教材处", "cyw1", "教材处", "教材处", "教材处", "教材处", "教材处", "教材处", "教材处", "教材处", "教材处"}
	var totalArr = []int{200, 400, 1, 2, 200, 1, 200, 2, 100, 80, 80, 230, 70, 700, 80}
	var collegeArr = []string{"北航学院", "计算机学院等", "北航学院", "北航学院", "北航学院", "计算机学院", "计算机学院", "经济管理学院", "法学院", "生物工程系", "网络空间安全学院", "计算机学院", "计算机学院", "北航学院", "计算机学院"}
	var priceArr = []int{68, 47, 35, 27, 18, 98, 34, 19, 18, 78, 23, 34, 24, 15, 24}
	var i int
	for i = 0; i < 15; i++ {
		textbook := new(deal.Textbook)
		textbook.BookName = bookNameArr[i]
		textbook.Writer = writerArr[i]
		textbook.Class = classArr[i]
		textbook.Description = descriptionArr[i]
		textbook.College = collegeArr[i]
		textbook.Price = int64(priceArr[i])
		textbook.Seller = sellerArr[i]
		textbook.Total = int64(totalArr[i])
		textbook.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		//deal.InsertTextbook(textbook)
	}
}
