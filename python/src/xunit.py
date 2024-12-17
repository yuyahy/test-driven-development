class TestResult:
    """テスト結果を定義するクラス"""

    def __init__(self):
        self.runCount = 0

    def testStarted(self):
        self.runCount = self.runCount + 1

    def summary(self):
        return "%d run, 0 failed" % self.runCount


class TestCase:
    """テストケースを定義するクラス"""

    def __init__(self, name):
        self.name = name

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def run(self):
        result = TestResult()
        result.testStarted()
        # fixture的な準備処理
        self.setUp()
        # コンストラクタで指定されたメソッド名を取得 & 実行
        method = getattr(self, self.name)
        method()
        # 後始末処理
        self.tearDown()
        return result


class WasRun(TestCase):
    def setUp(self):
        self.log = "setUp "

    def testMethod(self):
        self.log = self.log + "testMethod "

    def testBrokenMethod(self):
        raise Exception

    def tearDown(self):
        self.log = self.log + "tearDown "


class TestCaseTest(TestCase):
    """テストケースのテストを実行するクラス"""

    def testTemplateMethod(self):
        test = WasRun("testMethod")
        test.run()
        assert test.log == "setUp testMethod tearDown "

    def testResult(self):
        test = WasRun("testMethod")
        result = test.run()
        assert result.summary() == "1 run, 0 failed"

    def testFailedResult(self):
        test = WasRun("testBrokenMethod")
        result = test.run()
        assert result.summary() == "1 run, 1 failed"


def main():
    TestCaseTest("testTemplateMethod").run()
    TestCaseTest("testResult").run()
    # TestCaseTest("testFailedResult").run()


if __name__ == "__main__":
    main()
