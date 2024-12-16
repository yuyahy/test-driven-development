class TestCase:
    def __init__(self,name):
        self.name = name

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def run(self):
        # fixture的な準備処理
        self.setUp()
        # コンストラクタで指定されたメソッド名を取得 & 実行
        method = getattr(self,self.name)
        method()
        # 後始末処理
        self.tearDown()

class WasRun(TestCase):
    def setUp(self):
        self.log="setUp "

    def testMethod(self):
        self.log=self.log+"testMethod "

    def tearDown(self):
        self.log=self.log+"tearDown "

class TestCaseTest(TestCase):
    def testTemplateMethod(self):
        test=WasRun("testMethod")
        test.run()
        assert(test.log=="setUp testMethod tearDown ")

def main():
    #
    TestCaseTest("testTemplateMethod").run()

if __name__ == "__main__":
    main()